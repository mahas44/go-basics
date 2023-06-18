package channels

import "fmt"

func Select() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok bit", i, ok)
				return
			} else {
				fmt.Println("from comma ok bit", i, ok)
			}
		}
	}
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
