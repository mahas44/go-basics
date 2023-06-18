package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	shortLine := strings.Repeat("-", 20)
	fmt.Printf("| %s exercise-1 %s |\n", shortLine, shortLine)
	exercise1()
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
	fmt.Printf("| %s exercise-7 %s |\n", shortLine, shortLine)
	exercise7()
	fmt.Printf("| %s exercise-8 %s |\n", shortLine, shortLine)
	exercise8()
}

func exercise1() {
	a := foo1()
	b, c := bar1()
	fmt.Println(a, b, c)
}

func foo1() int {
	return 42
}

func bar1() (int, string) {
	return 451, "Fahrenheit"
}

func exercise2() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println("Foo2,", foo2(x...))
	fmt.Println("Bar2,", bar2(x))
}

func foo2(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar2(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func exercise3() {
	defer exercise2()
	fmt.Println("Bar3")
}

//exercise4
func exercise4() {
	p := person{
		first: "Enes",
		last:  "Çakmak",
		age:   26,
	}
	p.speak()
}

type person struct {
	first string
	last  string
	age   int16
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func exercise5() {
	c := circle{
		radius: 12.34,
	}
	s := square{
		edge: 15,
	}

	info(c)
	info(s)
}

type square struct {
	edge float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%T area: %f\n", s, s.area())
}

func (s square) area() float64 {
	return s.edge * s.edge
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func exercise6() {
	x := exercise5
	x()
}

func exercise7() {
	x := bar7()
	fmt.Println(x())
}

func bar7() func() int {
	return func() int {
		return 451
	}
}

func exercise8() {
	i := []int{100, 121, 144, 169}
	x := squareRoot(sum, i...)
	fmt.Println(x)
}

func squareRoot(f func(x ...int) int, params ...int) int {
	var y []int
	for _, v := range params {
		x := math.Sqrt(float64(v))
		y = append(y, int(x))
	}
	return sum(y...)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
