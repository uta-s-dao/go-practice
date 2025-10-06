package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func todo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todo.html")
	if err != nil {
		fmt.Println("cannot parse todofiles")
	}
	t.Execute(w, "Hello world!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", todo)
	fmt.Println("server on runnning:http://127.0.0.1:8080 ")
	server.ListenAndServe()
}
