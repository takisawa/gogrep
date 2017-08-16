package main

import (
	"regexp"
	"testing"
)

func TestMatchLine(t *testing.T) {
	var patternToLines = map[string]string{
		"text": "ABCtextXYZ",
		"日本語":  "あいう日本語わをん",
	}

	for patternText, line := range patternToLines {
		var pattern *regexp.Regexp = regexp.MustCompile(patternText)
		if !matchLine(pattern, line) {
			t.Errorf("`%s` is not match %s", line, pattern)
		}
	}
}
