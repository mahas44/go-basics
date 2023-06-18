package main

import "fmt"

func main() {
	// array()
	// slice()
	// appendElementIntoSlice()
	// deleteFromSlice()
	// makeSlice()
	// multiDimensionalSlice()
	maps()
}

func array() {
	var x [5]int
	fmt.Println(x)
	x[4] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}

func slice() {
	// COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	// slicing a slice
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:4])
	fmt.Println(x[:3])

	for i, v := range x[1:4] {
		fmt.Println(i, v)
	}
}

func appendElementIntoSlice() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 987}
	x = append(x, y...)
	fmt.Println(x)
}

func deleteFromSlice() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func makeSlice() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	//x[10] = 1234 // you can't do that // panic: runtime error: index out of range [10] with length 10
	x = append(x, 1234) // you can do that
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1235) // capacity full
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1236) // if you add new item in the full capacity slice, go runtime double the capacity
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multiDimensionalSlice() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Vanilla", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

func maps() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	// add item
	m["Todd"] = 34
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m)
	// delete item
	delete(m, "Todd")
	delete(m, "Barnabas")

	fmt.Println(m)

}
