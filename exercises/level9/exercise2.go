package level9

import "fmt"

/*
This exercise will reinforce our understanding of method sets:
	- create a type person struct
	- attach a method speak to type pointer to a person using a pointer receiver
		* *person
	- create a type human interface
		* to implicitly implement the interface, a human must have the speak method
	- create func "saySomething"
		* have it take in a human as a parameter
		* have it call the speak method
	- show the following in your code
		* you `can` pass a value of type *person into saySomething
		* you `cannot` pass a value of type person into saySomething
*/

type person struct {
	Name string
	Age  int16
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("My name is", p.Name, "and I am", p.Age)
}

func Exercise2() {
	p := person{
		Name: "James",
		Age:  32,
	}
	saySomething(&p)
}

func saySomething(h human) {
	h.speak()
}
