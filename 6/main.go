package main

//psql -h localhost -U gwp -d gwp

import (
	"database/sql"

	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB //グローバル変数の宣言

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=gwp dbname=gwp sslmode=disable") //変数への代入
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	err = rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content,author) values ($1,$2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Prepare failed:", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content=$2,author=$3 where id=$1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id=$1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Printf("GETPOST %+v\n", readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(100)
	fmt.Println(posts)

	readPost.Delete()

}

// func main() {
// 	data := []byte("Hello World\n")

// 	//WriteFileでファイルに書き込み、ReadFileでファイルから読み込む
// 	err := os.WriteFile("data1", data, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	read1, _ := os.ReadFile("data1")
// 	fmt.Print(string(read1))

// 	// 構造体Fileを使ってファイルの読み書きをする
// 	file1, _ := os.Create("data2")
// 	defer file1.Close()

// 	bytes, _ := file1.Write(data)
// 	fmt.Printf("write %d bytes to file\n", bytes)

// 	file2, _ := os.Open("data2")
// 	defer file2.Close()

// 	read2 := make([]byte, len(data))
// 	bytes, _ = file2.Read(read2)
// 	fmt.Printf("read %d bytes from file\n", bytes)
// 	fmt.Print(string(read2))
// }
