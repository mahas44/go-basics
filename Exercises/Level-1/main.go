package main

import (
	"fmt"
	"strings"
)

var x int
var y string
var z bool

type mytype int

var q mytype
var w int

func main(){
	shortLine:= strings.Repeat("-", 20)
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
}

func exercise1(){
	x:= 42
	y:="James Bond"
	z:= true

	fmt.Println(x,y,z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise2(){
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise3() {
	x = 42
	y = "James Bond"
	z = true;

	s:= fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

func exercise4() {
	fmt.Println(q)
	fmt.Printf("%T\n", q)
	q = 42
	fmt.Println(q)
}

func exercise5() {
	fmt.Println(q)
	fmt.Printf("%T\n", q)
	q = 42
	fmt.Println(q)
	w = int(q)
	fmt.Println(w)
	fmt.Printf("%T\n", w)
}