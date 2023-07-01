package errorhandling

import (
	"fmt"
	"log"
	"os"
)

func Logging() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// comment row and run
		fmt.Println("error happened", err)
		log.Println("error happened", err)
		log.Fatalln(err)
		log.Panic(err)
		panic(err)
	}
}
