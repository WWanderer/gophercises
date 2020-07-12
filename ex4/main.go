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

	t, err := linkextractor.GetHTMLTree(f)
	if err != nil {
		panic(err)
	}

	links := linkextractor.ExtractLinks(t)
	fmt.Println(links)
}
