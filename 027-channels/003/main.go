package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	q := make(chan int)

	//send
	go send(even, odd, q)

	//receive
	receive(even, odd, q)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Received from even channel:", v)
		case v := <-o:
			fmt.Println("Received from odd channel", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("Received from quit channel", v)
				return
			}
		}
	}
}
