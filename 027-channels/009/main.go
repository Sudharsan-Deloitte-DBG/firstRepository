package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			select {
			case <-q:
				return
			default:
				c <- i
			}
		}
	}()

	return c
}

func receive(c <-chan int, q <-chan int) {
	select {
	case <-q:
		return
	default:
		for v := range c {
			fmt.Println(v)
		}
	}
}
