package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	d1 := time.Duration(rand.Int63n(100))
	d2 := time.Duration(rand.Int63n(200))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		c1 <- 500
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		c2 <- 400
	}()

	select {
	case v1 := <-c1:
		fmt.Println("Channel 1 wins: ", v1)
	case v2 := <-c2:
		fmt.Println("Channel 2 wins: ", v2)
	}
}
