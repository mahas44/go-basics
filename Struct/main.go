package main

import "fmt"

type person struct {
	first string
	last  string
	age   int16
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneyponny",
		age:   27,
	}

	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Ian",
			last:  "Felming",
			age:   42,
		},
		ltk: true,
	}

	fmt.Println(p1, p1.first, p1.last, p1.age)
	fmt.Println(p2, p2.first, p2.last, p2.age)

	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa2)

	anonymousStruct()
}

func anonymousStruct() {
	p1 := struct {
		first string
		last  string
		age   int16
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println("anonymousStruct ", p1)
}
