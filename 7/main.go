package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"os"
// )

// type Post struct {
// 	XMLName xml.Name `xml:"post"`
// 	Id      string   `xml:"id,attr"`
// 	Content string   `xml:"content"`
// 	Author  Author   `xml:"author"`
// }

// type Author struct {
// 	Id   string `xml:"id,attr"`
// 	Name string `xml:"name"`
// }

// func main() {
// 	post := Post{
// 		XMLName: xml.Name{Local: "post"},
// 		Id:      "1",
// 		Content: "Hello World!",
// 		Author: Author{
// 			Id:   "2",
// 			Name: "Sau Sheong",
// 		},
// 	}

// 	output, err := xml.Marshal(&post)
// 	if err != nil {
// 		fmt.Println("Error creating XML file:", err)
// 		return
// 	}
// 	err = os.WriteFile("post.xml", output, 0644)
// 	if err != nil {
// 		fmt.Println("Error encoding XML to file:", err)
// 		return
// 	}
// }
