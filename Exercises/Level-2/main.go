package main

import (
	"fmt"
	"strings"
)

const (
	q int = 42
	w     = "James Bond"
)

const (
	p = 2019 + iota
	r = 2019 + iota
	s = 2019 + iota
	t = 2019 + iota
)

func main() {
	shortLine := strings.Repeat("-", 20)
	fmt.Printf("| %s exercise-2 %s |\n", shortLine, shortLine)
	exercise2()
	fmt.Printf("| %s exercise-3 %s |\n", shortLine, shortLine)
	exercise3()
	fmt.Printf("| %s exercise-4 %s |\n", shortLine, shortLine)
	exercise4()
	fmt.Printf("| %s exercise-5 %s |\n", shortLine, shortLine)
	exercise5()
	fmt.Printf("| %s exercise-6 %s |\n", shortLine, shortLine)
	exercise6()
}

func exercise1() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}

func exercise2() {
	a := 42 == 42
	b := 42 <= 43
	c := 42 >= 41
	d := 42 != 42
	e := 42 > 40
	f := 42 < 43
	fmt.Println(a, b, c, d, e, f)
}

func exercise3() {
	fmt.Printf("%d\t\t%s\n", q, w)
}

func exercise4() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	a = a << 1
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}

func exercise5() {
	a := `Here is something
	as
	a
	raw
	string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}

func exercise6() {
	fmt.Printf("%d\t%d\t%d\t%d\n", p, r, s, t)
}
