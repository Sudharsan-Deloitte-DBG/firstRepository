package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	wg.Add(2)
	go foo()
	go bar()
	ch := make(chan int)
	go func(i int) {
		ch <- DoSomething(i)
	}(10)

	fmt.Println("Value from channel:", <-ch)

	fmt.Println("Go Routines", runtime.NumGoroutine())
	fmt.Println("Go Routines", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}

func DoSomething(i int) int {
	for ; i > 20; i++ {
		fmt.Println("Incremented i to", i)
	}
	return i
}
