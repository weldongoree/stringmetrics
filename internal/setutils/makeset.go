package setutils

import (
	"strings"
	"unicode/utf8"
)

func makeRuneSet(input string) map[rune]bool {
	output := make(map[rune]bool)
	for len(input) > 0 {
		r, size := utf8.DecodeLastRuneInString(input)
		output[r] = true
		input = input[:len(input)-size]
	}
	return output
}

func makeWordSet(input string) map[string]bool {
	output := make(map[string]bool)
	words := strings.Fields(input)
	for _, w := range words {
		output[w] = true
	}
	return output
}
