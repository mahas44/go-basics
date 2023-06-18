package functions

import (
	"fmt"
)

func Recursion() {
	x := factoriel(4)
	fmt.Println(x)
}

func factoriel(n int) int {
	if n == 0 {
		return 1
	}
	return n * factoriel(n-1)
}
