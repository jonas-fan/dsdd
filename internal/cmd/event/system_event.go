package event

import (
	"fmt"
	"os"
)

func readSystemEvent(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println(filename)
}
