package main

import (
	"fmt"
)

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	ok1, ok2 := true, true

	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from a\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from b\n", msg)
			}
		}
	}
}

// import (
// 	"fmt"
// 	"time"
// )

// func thrower(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 		fmt.Println("Threw >>", i)
// 	}
// }

// func catcher(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		num := <-c
// 		fmt.Println("Caught <<", num)
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	go thrower(c)
// 	go catcher(c)
// 	time.Sleep(100 * time.Millisecond)
// }
