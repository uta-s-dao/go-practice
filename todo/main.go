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
	list, err := getAll()
	if err != nil {
		fmt.Println("cannnot get todolist")
	}
	fmt.Println(list)
	if err != nil {
		fmt.Println("cannot marshal struct to json")
	}
	t.Execute(w, list)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/todo", todo)
	http.HandleFunc("/post/", handleRequest)
	fmt.Println("server on runnning:http://127.0.0.1:8080/todo/ and http://127.0.0.1:8080/post/")
	server.ListenAndServe()
}

//curl -i -X POST -H "Content-Type: application/json" -d '{"name":"breakfast","status":"todo"}' http://127.0.0.1:8080/post/
