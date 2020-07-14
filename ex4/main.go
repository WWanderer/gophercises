package main

import (
	"fmt"
	"os"

	"./linkextractor"
)

func main() {
	f, err := os.Open("./ex2.html")
	if err != nil {
		panic(err)
	}

	fmt.Println(linkextractor.Parse(f))
}
