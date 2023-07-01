package level11

import "fmt"

type customError struct {
	info string
}

func (c customError) Error() string {
	return fmt.Sprintf("error: %v", c.info)
}

func Exercise3() {
	c1 := customError{info: "need more tea"}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n", e.(customError).info)
}
