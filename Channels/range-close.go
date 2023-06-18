package channels

import "fmt"

func RangeClose() {
	c := make(chan int)
	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	// recive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
