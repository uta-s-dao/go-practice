package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello World\n")

	//WriteFileでファイルに書き込み、ReadFileでファイルから読み込む
	err := os.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := os.ReadFile("data1")
	fmt.Print(string(read1))

	// 構造体Fileを使ってファイルの読み書きをする
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("write %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("read %d bytes from file\n", bytes)
	fmt.Print(string(read2))
}
