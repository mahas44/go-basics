package level11

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

var p1 = person{
	First:   "James",
	Last:    "Bond",
	Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
}

func Exercise1() {
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal: ", err)
	}
	fmt.Println(string(bs))
}
