package main

import (
	"fmt"
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
	fmt.Printf("| %s exercise-9 %s |\n", shortLine, shortLine)
	exercise9()
	fmt.Printf("| %s exercise-10 %s |\n", shortLine, shortLine)
	exercise10()
}

func exercise1() {
	x := [5]int{42, 43, 44, 45, 46}
	fmt.Printf("x: %v\n", x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}

func exercise2() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("x: %v\n", x)
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}

func exercise3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exercise6() {
	x := make([]string, 0, 12)
	y := []string{"Adana", "Antalya", "Ankara", "Bursa", "Çanakkale", "Çorum", "Denizli", "Edirne", "Hatay", "Isparta", "İstanbul", "Malatya"}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("len: ", len(x))
	fmt.Println("cap: ", cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func exercise7() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneyponny", "Hello James"}

	z := [][]string{x, y}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println("RECORD", i, v)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v \t value: %v\n", j, v2)
		}
	}
}

func exercise8() {
	x := map[string][]string{}
	james := []string{"Shaken, not stirred", "Martinis", "Women"}
	moneypenny := []string{"James Bond", "Literature", "Computer Science"}
	dr := []string{"Being evil", "Ice cream", "Sunsets"}
	x["bond_james"] = james
	x["moneyponny_miss"] = moneypenny
	x["no_dr"] = dr

	for i, v := range x {
		fmt.Println("RECORD", i)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v \t value: %v\n", j, v2)
		}
	}
}

func exercise9() {
	x := map[string][]string{}
	james := []string{"Shaken, not stirred", "Martinis", "Women"}
	moneypenny := []string{"James Bond", "Literature", "Computer Science"}
	dr := []string{"Being evil", "Ice cream", "Sunsets"}
	x["bond_james"] = james
	x["moneyponny_miss"] = moneypenny
	x["no_dr"] = dr

	x["fleming_ian"] = []string{"Steaks", "Cigars", "Espionage"}

	for i, v := range x {
		fmt.Println("RECORD", i)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v \t value: %v\n", j, v2)
		}
	}
}

func exercise10() {
	x := map[string][]string{}
	james := []string{"Shaken, not stirred", "Martinis", "Women"}
	moneypenny := []string{"James Bond", "Literature", "Computer Science"}
	dr := []string{"Being evil", "Ice cream", "Sunsets"}
	x["bond_james"] = james
	x["moneyponny_miss"] = moneypenny
	x["no_dr"] = dr

	x["fleming_ian"] = []string{"Steaks", "Cigars", "Espionage"}

	delete(x, "no_dr")

	for i, v := range x {
		fmt.Println("RECORD", i)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v \t value: %v\n", j, v2)
		}
	}
}
