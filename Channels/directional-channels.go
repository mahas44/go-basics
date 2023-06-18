package channels

import "fmt"

func DreictionalChannels() {
	c := make(chan int, 2)
	// read from left to right

	// receive: can read from, but cannot write to or close()
	// <-chan // data only comes out
	cr := make(<-chan int, 2)
	// send: cannot read from, but can write to and close()
	// chan<- // data only goes in
	cs := make(chan<- int, 2)

	// specific to general does not convert exp: c = cr
	// general to specific can be convert exp: cr = c
	fmt.Println("------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
