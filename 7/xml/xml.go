package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Postun struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Authorm  `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Authorm struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name"`
}

func main() {
	xmlFile, err := os.Open("postun.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var postun Postun
	err = xml.Unmarshal(xmlData, &postun)
	if err != nil {
		fmt.Println("Error unmarshalling XML data:", err)
		return
	}
	fmt.Println(postun)
}
