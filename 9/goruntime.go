package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		// fmt.Printf("%d", i)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		// fmt.Printf("%c", i)
	}
}

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c", i)
	}
	w <- true
}

// func print1() {
// 	printNumbers()
// 	printLetters()
// }

// func print2() {
// 	printNumbers2(wg * sync.WaitGroup)
// 	printLetters2()
// }

// func goPrint1() {
// 	go printNumbers()
// 	go printLetters()
// }

// func goPrint2() {
// 	go printNumbers2()
// 	go printLetters2()
// }

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers2(w1)
	go printLetters2(w2)
	<-w1
	<-w2
}
