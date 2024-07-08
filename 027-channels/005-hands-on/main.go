package main

import "fmt"

func main() {
	c := make(chan int)
	d := make(chan int, 2)
	d <- 20
	go func() {
		c <- 10
	}()

	fmt.Println(<-c)
	fmt.Println(<-d)
}
