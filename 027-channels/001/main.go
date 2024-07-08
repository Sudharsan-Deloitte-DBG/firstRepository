package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 100
	c <- 200

	fmt.Println(<-c)

	roc := make(chan<- int)
	woc := make(<-chan int)

	go func() {
		roc <- 10
		<-woc
	}()
}
