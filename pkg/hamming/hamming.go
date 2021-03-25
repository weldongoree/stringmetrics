package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(str1 string, str2 string) (int, error) {
	output := 0
	if len(str1) != len(str2) {
		err := errors.New("incompatible string lengths")
		return 0, err
	}
	for len(str1) > 0 {
		r1, size1 := utf8.DecodeLastRuneInString(str1)
		r2, size2 := utf8.DecodeLastRuneInString(str2)
		if r1 != r2 {
			output += 1
		}
		str1 = str1[:len(str1)-size1]
		str2 = str2[:len(str2)-size2]
		// This worries me, but let's assume they are normalized for now
	}

	return output, nil
}
