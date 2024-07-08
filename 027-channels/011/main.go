package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- rand.Intn(500000)
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
