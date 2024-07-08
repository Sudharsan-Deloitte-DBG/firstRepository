package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	ngr := 100
	var wg sync.WaitGroup
	wg.Add(ngr)

	//var mut sync.Mutex
	var counter int64
	for i := 0; i < ngr; i++ {
		go func() {
			//mut.Lock()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			//mut.Unlock()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Num of goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
