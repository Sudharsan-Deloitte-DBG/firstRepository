package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

//var mut sync.Mutex

func main() {
	var counter int32
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			//mut.Lock()
			atomic.AddInt32(&counter, 1)

			//mut.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
