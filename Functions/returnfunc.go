package functions

import (
	"fmt"
)

func ReturnFunc() {
	x := bar2()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
	fmt.Println(x())
	fmt.Println(bar2()())

}

func bar2() func() int {
	return func() int {
		return 451
	}
}

// callback func
func sum2(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
