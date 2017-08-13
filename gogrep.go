package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var args []string
	var regexp_text string
	var pattern *regexp.Regexp
	var done = make(chan struct{})

	var caseInsensitive = flag.Bool("i", false, "case-insensitive")

	flag.Parse()

	args = flag.Args()

	if flag.NArg() < 2 {
		fmt.Println("Invalid arguments")
		os.Exit(1)
	}

	regexp_text = args[0]

	if *caseInsensitive {
		regexp_text = "(?i)" + regexp_text
	}

	pattern = regexp.MustCompile(regexp_text)

	for _, fname := range args[1:] {
		go grepOneFile(os.Stdout, pattern, fname, done)
	}

	for i := 0; i < flag.NArg()-1; i++ {
		<-done
	}

	close(done)
}
