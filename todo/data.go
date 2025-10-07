package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type todoList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=todo password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (todos todoList, err error) {
	todos = todoList{}
	err = Db.QueryRow("select id, name, status from todos where id = $1", id).Scan(&todos.Id, &todos.Name, &todos.Status)
	return
}

func getAll() (allTodo []todoList, err error) {
	rows, err := Db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allTodo = []todoList{}
	for rows.Next() {
		var todo todoList
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Status)
		if err != nil {
			return nil, err
		}
		allTodo = append(allTodo, todo)
	}
	return allTodo, nil
}

func (todos *todoList) Create() (err error) {
	statement := "insert into todos (name,status) values ($1,$2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(todos.Name, todos.Status).Scan(&todos.Id)
	return
}

func (todos *todoList) Update() (err error) {
	_, err = Db.Exec("update todos set name=$2,status=$3 where id=$1", todos.Id, todos.Name, todos.Status)
	return
}

func (todos *todoList) Delete() (err error) {
	_, err = Db.Exec("delete from todos where id=$1", todos.Id)
	return
}
