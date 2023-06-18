package functions

import (
	"fmt"
)

func Closure() {
	//firstExp()
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func firstExp() {
	var x int
	fmt.Println(x)
	x++

	{
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y) //it is not work. because y is defined in the block
	fmt.Println(x)

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
