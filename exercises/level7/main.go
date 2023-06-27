package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p := person{
		firstName: "James",
		lastName:  "Bond",
	}
	fmt.Println(p)
	fmt.Println(&p.firstName)
	changeMe(&p)
	fmt.Println(p)
	fmt.Println(&p.firstName)
}

func changeMe(p *person) {
	p.firstName = "Miss"
	(*p).lastName = "Moneypenny"
}
