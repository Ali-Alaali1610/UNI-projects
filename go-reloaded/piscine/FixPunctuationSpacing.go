package piscine

import (
	"unicode"
)

func FixPunctuationSpacing(sentence string) string {
	var output string

	for i := 0; i < len(sentence); i++ {
		output += string(sentence[i])

		if unicode.IsPunct(rune(sentence[i])) &&
			i+1 < len(sentence) && !unicode.IsPunct(rune(sentence[i+1])) {

			output += " "
		}
	}
	return output
}
