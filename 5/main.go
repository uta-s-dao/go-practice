package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatedDate(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html", "index1.html")
	t.Execute(w, "hello world")
}

func date(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatedDate}
	t := template.New("date.html").Funcs(funcMap)
	t.ParseFiles("date.html")
	t.Execute(w, time.Now())
}

func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/date", date)
	server.ListenAndServe()
}
