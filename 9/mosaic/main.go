package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.↓")
	fmt.Println("http://127.0.0.1:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	err := r.ParseMultipartForm(10485760)
	if err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get image file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	tileSize, err := strconv.Atoi(r.FormValue("tile_size"))
	if err != nil {
		http.Error(w, "Invalid tile size: "+err.Error(), http.StatusBadRequest)
		return
	}

	original, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image: "+err.Error(), http.StatusBadRequest)
		return
	}
	bounds := original.Bounds()

	newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	db := cloneTilesDB()
	fmt.Printf("Cloned DB has %d tiles\n", len(db))

	sp := image.Point{0, 0}
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {

			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}

			nearest := nearest(color, &db)
			if nearest == "" {
				fmt.Println("Warning: no more tiles available in database")
				continue
			}
			file, err := os.Open(nearest)
			if err != nil {
				fmt.Printf("error opening file: %v, filename: '%s'\n", err, nearest)
				continue
			}

			img, _, err := image.Decode(file)
			if err != nil {
				fmt.Println("error decoding image:", err, nearest)
				file.Close()
				continue
			}

			t := resize(img, tileSize)
			tile := t.SubImage(t.Bounds())
			tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)

			draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newimage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)

}
