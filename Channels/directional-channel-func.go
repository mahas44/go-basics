package channels

import "fmt"

func DirectionalChannelFunc() {
	c := make(chan int)
	// send
	go foo(c)
	// recive
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
