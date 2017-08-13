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

func grepOneFile(out *os.File, pattern *regexp.Regexp, fname string, done chan struct{}) {
	var fp *os.File
	var err error
	var reader *bufio.Reader

	fp, err = os.Open(fname)

	if err != nil {
		panic(err)
	}

	defer fp.Close()

	reader = bufio.NewReaderSize(fp, BUFSIZE)
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		if pattern.MatchString(line) {
			fmt.Fprint(out, line)
		}
	}

	if err != io.EOF {
		panic(err)
	}

	done <- struct{}{}
}

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
