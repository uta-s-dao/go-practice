package main

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type todoList struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func todo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todo.html")
	if err != nil {
		fmt.Println("cannot parse todofiles")
	}
	list := []todoList{{Name: "guiter", Status: "complete"}, {Name: "dinner", Status: "todo"}}
	// jsondata, err := json.MarshalIndent(list, "", "")
	fmt.Println(list)
	if err != nil {
		fmt.Println("cannot marshal struct to json")
	}
	// fmt.Println(string(jsondata))
	t.Execute(w, list)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/todo", todo)
	fmt.Println("server on runnning:http://127.0.0.1:8080/todo")
	server.ListenAndServe()
}
