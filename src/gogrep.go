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

func grepOneFile(pattern *regexp.Regexp, fname string) {
	var fp *os.File
	var err error

	fp, err = os.Open(fname)

	if err != nil {
		panic(err)
	}

	defer fp.Close()

	reader := bufio.NewReaderSize(fp, BUFSIZE)
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		if pattern.MatchString(line) {
			fmt.Print(line)
		}
	}

	if err != io.EOF {
		panic(err)
	}
}

func main() {
	var caseInsensitive = flag.Bool("i", false, "case-insensitive")
	var pattern *regexp.Regexp

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

	pattern = regexp.MustCompile(regexp_text)

	for _, fname := range args[1:] {
		grepOneFile(pattern, fname)
	}
}
