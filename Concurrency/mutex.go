package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func MutexExp() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mtx.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

	// command: go run -race mutex.go
}
