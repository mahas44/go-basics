package level9

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	- Using goroutines, create an incrementer program
		- have a variable to hold the incrementer value
		- launch a bunch of goroutines
			- each goroutine should
				- read the incrementer value
					- store it in a new variable
				- yield the processor with runtime.Gosched()
				- increment the new variable
				- write the value in the new variable back to the incrementer variable
	- use waitgroups to wait for all of your goruputine to finish
	- the above will create a race condition
	- prove that it is a race condition by using the -race flag
*/

func Exercise3() {

	counter := 0
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
