package errorhandling

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func ErrorInfo() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Printf("%T\n", ErrNorgateMath)
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//ErrNorgateMath = fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	return 42, nil
}
