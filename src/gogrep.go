package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Invalid arguments")
		os.Exit(1)
	}

	regexp_text := os.Args[1]
	files := os.Args[2:]

	fmt.Printf("regexp_text: %s\n", regexp_text)

	for i, file := range files {
		fmt.Printf("file: %d: %s\n", i, file)
	}
}
