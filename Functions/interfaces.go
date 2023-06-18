package functions

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

type salary interface {
	getSalary(s string) float64
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func (sa secretAgent) getSalary(s string) float64 {
	result := 0.0
	switch s {
	case "James":
		result = 10000.0
	case "Miss":
		result = 8000.0
	}
	return result
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first,
			"and this is my salary:", h.(secretAgent).getSalary(h.(secretAgent).first), "$")
	default:
		fmt.Println("Unknown type", h)
	}
}

func Interfaces() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr",
		last:  "Yes",
	}

	bar(sa1)
	bar(sa2)
	bar(p1)
}
