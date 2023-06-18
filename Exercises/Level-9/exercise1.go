package exercises

import (
	"fmt"
	"runtime"
	"sync"
)

/*
in addition to the main gorountine, launch two additional gorountines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/

func Exercise1() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("foo")
		wg.Done()
	}()

	go func() {
		fmt.Println("bar")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
