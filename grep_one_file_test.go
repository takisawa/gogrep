package main

import (
	"regexp"
	"testing"
)

func TestMatchLine(t *testing.T) {
	var patternText = "text"
	var pattern *regexp.Regexp = regexp.MustCompile(patternText)
	var line string = "ABCtextXYZ"

	if !matchLine(pattern, line) {
		t.Errorf("`%s` is not match %s", line, pattern)
	}
}
