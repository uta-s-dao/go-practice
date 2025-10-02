package main

// import (
// 	"fmt"
// 	"image"
// 	_ "image/gif"
// 	"image/jpeg"
// 	_ "image/png"
// 	"os"
// 	"path/filepath"
// )

// func Generete() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("Usage: go run generate_tiles.go <input_image> <tile_size>")
// 		fmt.Println("Example: go run generate_tiles.go source.jpg 50")
// 		return
// 	}

// 	inputPath := os.Args[1]
// 	tileSize := 50
// 	fmt.Sscanf(os.Args[2], "%d", &tileSize)

// 	// 入力画像を開く
// 	file, err := os.Open(inputPath)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// 画像をデコード
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		fmt.Println("Error decoding image:", err)
// 		return
// 	}

// 	bounds := img.Bounds()
// 	width := bounds.Max.X - bounds.Min.X
// 	height := bounds.Max.Y - bounds.Min.Y

// 	fmt.Printf("Image size: %dx%d\n", width, height)
// 	fmt.Printf("Tile size: %d\n", tileSize)

// 	tileCount := 0
// 	// 画像を分割
// 	for y := 0; y < height-tileSize; y += tileSize {
// 		for x := 0; x < width-tileSize; x += tileSize {
// 			// タイル領域を切り出す
// 			tileRect := image.Rect(x, y, x+tileSize, y+tileSize)
// 			tile := img.(interface {
// 				SubImage(r image.Rectangle) image.Image
// 			}).SubImage(tileRect)

// 			// タイルを保存
// 			outputPath := filepath.Join("tiles", fmt.Sprintf("tile_%04d.jpg", tileCount))
// 			outFile, err := os.Create(outputPath)
// 			if err != nil {
// 				fmt.Println("Error creating tile file:", err)
// 				continue
// 			}

// 			err = jpeg.Encode(outFile, tile, &jpeg.Options{Quality: 75})
// 			if err != nil {
// 				fmt.Println("Error encoding tile:", err)
// 			}
// 			outFile.Close()
// 			tileCount++
// 		}
// 	}

// 	fmt.Printf("Generated %d tiles in tiles/ directory\n", tileCount)
// }

// func main() {
// 	Generete()
// }
