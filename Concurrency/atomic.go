package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func AtomicExp() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

	// command: go run -race mutex.go
}
