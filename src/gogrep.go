package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const BUFSIZE = 4096

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 3 {
		fmt.Println("Invalid arguments")
		os.Exit(1)
	}

	regexp_text := os.Args[1]
	regexp := regexp.MustCompile(regexp_text)

	files := os.Args[2:]

	for _, file := range files {
		fp, err = os.Open(file)

		if err != nil {
			panic(err)
		}

		defer fp.Close()

		reader := bufio.NewReaderSize(fp, BUFSIZE)
		for line := ""; err == nil; line, err = reader.ReadString('\n') {
			if regexp.MatchString(line) {
				fmt.Print(line)
			}
		}

		if err != io.EOF {
			panic(err)
		}
	}
}
