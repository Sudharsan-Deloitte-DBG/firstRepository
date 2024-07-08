package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go addNumsToChannel(ch)
	}

	for j := 0; j < 100; j++ {
		fmt.Println(<-ch)
	}
}

func addNumsToChannel(respch chan<- int) {
	for i := 0; i < 10; i++ {
		respch <- rand.Intn(50000)
	}
}
