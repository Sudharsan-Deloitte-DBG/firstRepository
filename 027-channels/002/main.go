package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 2)

	//send
	go foo(c)

	//receive
	go bar(c)

	time.Sleep(2 * time.Millisecond)

	fmt.Println("From main goroutine:", <-c)

	fmt.Println("About to exit!!")
}

// send
func foo(c chan<- int) {
	c <- 42
	c <- 43
}

// receive
func bar(c <-chan int) {
	fmt.Println("From bar goroutine:", <-c)
}
