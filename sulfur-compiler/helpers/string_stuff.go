package helpers

import (
	"regexp"
	"strings"
)

// from: https://github.com/lithammer/dedent/blob/master/dedent.go
// MIT license

var (
	leadingWhitespace = regexp.MustCompile("(?m)(^[ \t]*)(?:[^ \t\n])")
)

// also removes leading newline (if any)
// also replaces tabs with four spaces
// trims any trailing spaces (but leaves newlines)
func Dedent(text string) string {
	var margin string

	tabless := strings.Replace(text, "\t", "    ", -1)
	tabless = strings.TrimLeft(tabless, "\n")
	tabless = strings.TrimRight(tabless, " ")
	indents := leadingWhitespace.FindAllStringSubmatch(tabless, -1)

	// Look for the longest leading string of spaces and tabs common to all
	// lines.
	for i, indent := range indents {
		if i == 0 {
			margin = indent[1]
		} else if strings.HasPrefix(indent[1], margin) {
			// Current line more deeply indented than previous winner:
			// no change (previous winner is still on top).
			continue
		} else if strings.HasPrefix(margin, indent[1]) {
			// Current line consistent with and no deeper than previous winner:
			// it's the new winner.
			margin = indent[1]
		} else {
			// Current line and previous winner have no common whitespace:
			// there is no margin.
			margin = ""
			break
		}
	}

	if margin != "" {
		tabless = regexp.MustCompile("(?m)^"+margin).ReplaceAllString(tabless, "")
	}
	return tabless
}
