package level10

import "fmt"

func Exercise1() {
	// error
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)
	sol1()
	sol2()
}

func sol1() {
	c := make(chan int)
	// anonymous self-executing fun
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

func sol2() {
	// buffered channel
	c := make(chan int, 1)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
