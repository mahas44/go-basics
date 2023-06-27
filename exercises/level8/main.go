package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type user struct {
	First string
	Age   int
}

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []person

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	shortLine := strings.Repeat("-", 20)
	// fmt.Printf("| %s exercise-1 %s |\n", shortLine, shortLine)
	// exercise1()
	// fmt.Printf("| %s exercise-2 %s |\n", shortLine, shortLine)
	// exercise2()
	// fmt.Printf("| %s exercise-3 %s |\n", shortLine, shortLine)
	// exercise3()
	fmt.Printf("| %s exercise-4 %s |\n", shortLine, shortLine)
	exercise4()
}

func exercise1() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypeny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}
	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(bs))
}

func exercise2() {

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs := []byte(s)
	var people []person
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, val := range person.Sayings {
			fmt.Println("\t\t", val)
		}
	}
}

func exercise3() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	p3 := person{
		First: "M",
		Last:  "Hmmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []person{p1, p2, p3}
	fmt.Println(people)

	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}

func exercise4() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	p3 := person{
		First: "M",
		Last:  "Hmmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	people := []person{p1, p2, p3}
	fmt.Println(people)

	for _, p := range people {
		fmt.Println(p.First, p.Last, p.Age)
		sort.Strings(p.Sayings)
		for _, v := range p.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("---------------------")
	sort.Sort(ByAge(people))
	for _, p := range people {
		fmt.Println(p.First, p.Last, p.Age)
		for _, v := range p.Sayings {
			fmt.Println("\t", v)
		}
	}
	fmt.Println("---------------------")
	sort.Sort(ByLast(people))
	for _, p := range people {
		fmt.Println(p.First, p.Last, p.Age)
		for _, v := range p.Sayings {
			fmt.Println("\t", v)
		}
	}
}
