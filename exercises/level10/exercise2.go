package level10

import "fmt"

func Exercise2() {
	// cs := make(chan<- int)
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)
	fmt.Println("--------")
	fmt.Printf("cs\t%T\n", cs)
}
