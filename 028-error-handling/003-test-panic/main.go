package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	defer func() {
		fmt.Println("Deferred function in main goroutine")
	}()
	fmt.Println("Inside main goroutine")
	var wg sync.WaitGroup
	wg.Add(2)
	go panicTest(&wg)
	go noPanic(&wg)
	wg.Wait()
	fmt.Println("Test to see if panic unwinds the call stack of a single goroutine or unwinds call stacks of all goroutines.")

}

func noPanic(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("This is a non panicking goroutine created from main")
}

func panicTest(wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Deferred function in panicTest goroutine")
		wg.Done()
	}()
	fmt.Println("Inside panicTest goroutine")
	time.Sleep(time.Millisecond * 50)
	panicTestLayer2()
}

func panicTestLayer2() {
	defer func() {
		fmt.Println("Deferred function in panicTestLayer2")
	}()
	fmt.Println("Inside panicTestLayer2")
	time.Sleep(time.Millisecond * 200)
	panicTestLayer3()
}

func panicTestLayer3() {
	defer func() {
		fmt.Println("Deferred function in panicTestLayer3")
	}()
	fmt.Println("Inside panicTestLayer3")
	time.Sleep(time.Millisecond * 200)

	_, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}

}
