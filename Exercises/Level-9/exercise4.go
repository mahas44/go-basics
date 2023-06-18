package exercises

/*
	Exercise3 with mutex
*/
import (
	"fmt"
	"runtime"
	"sync"
)

func Exercise4() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mtx.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
