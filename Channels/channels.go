package channels

import "fmt"

func ChannelExp() {
	// does not work
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)

	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println("Normal channel", <-c)

	// buffer channel
	bc := make(chan int, 1)
	bc <- 42
	fmt.Println("Buffer channel", <-bc)

	// does not work - deadlock
	// bc2 := make(chan int, 1)
	// bc2 <- 42
	// bc2 <- 43
	// fmt.Println("Buffer channel", <-bc2)

	// buffer channel
	bc2 := make(chan int, 2)
	bc2 <- 42
	bc2 <- 43
	fmt.Println("Buffer channel2", <-bc2)
	fmt.Println("Buffer channel2", <-bc2)
}
