package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

const BUFSIZE = 4096

func main() {
	var fp *os.File
	var err error
	var caseInsensitive = flag.Bool("i", false, "case-insensitive")

	flag.Parse()

	var args = flag.Args()

	if len(os.Args) < 3 {
		fmt.Println("Invalid arguments")
		os.Exit(1)
	}

	regexp_text := args[0]

  if *caseInsensitive {
    regexp_text = "(?i)" + regexp_text
  }

	regexp := regexp.MustCompile(regexp_text)

	for _, file := range args[1:] {
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
