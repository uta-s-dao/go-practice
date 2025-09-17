package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

func main() {

	post := Post{
		Id:      1,
		Content: "Hello World",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id:      1,
				Content: "Good post",
				Author: Author{
					Id:   3,
					Name: "Joe",
				},
			},
			{
				Id:      2,
				Content: "Thanks for the info",
				Author: Author{
					Id:   4,
					Name: "Mary",
				},
			},
		},
	}

	output, err := json.MarshalIndent(post, "", "\t\t ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = os.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	// jsonFile, err := os.Create("post.json")
	// if err != nil {
	// 	fmt.Println("Error creating JSON file:", err)
	// 	return
	// }

	// encoder := json.NewEncoder(jsonFile) //エンコーダー作成
	// err = encoder.Encode(post)           //構造体 → JSON変換 → ファイル書き込み
	// if err != nil {
	// 	fmt.Println("Error encoding JSON to file:", err)
	// 	return
	// }

}
