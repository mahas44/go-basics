package level11

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (s sqrtError) Error() string {
	return fmt.Sprintf("math error: lat: %v, long: %v, err: %v", s.lat, s.long, s.err)
}

func Exercise4() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			lat:  "55.1234 N",
			long: "99.3432 W",
			err:  fmt.Errorf("square root of negative number -> %v", f),
		}
	}
	return 42, nil
}
