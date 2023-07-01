package level11

import (
	"encoding/json"
	"fmt"
	"log"
)

func Exercise2() {
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

// toJSON needs to return an error
func toJSON(a interface{}) ([]byte, error) {

	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in to json: %v ", err)
	}
	return bs, nil
}
