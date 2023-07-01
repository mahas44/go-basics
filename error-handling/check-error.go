package errorhandling

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Check() {
	one()
	two()
	three()
	four()
}

func one() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func two() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
}

func three() {
	f, err := os.Create("./error-handling/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}

func four() {
	f, err := os.Open("./error-handling/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
