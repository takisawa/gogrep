package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const BUFSIZE = 4096

func matchLine(pattern *regexp.Regexp, line string) bool {
	return pattern.MatchString(line)
}

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
		if matchLine(pattern, line) {
			fmt.Fprint(out, line)
		}
	}

	if err != io.EOF {
		panic(err)
	}

	done <- struct{}{}
}
