package piscine

import (
	"regexp"
	"strings"
)

func UpperWord(words []string) []string {
	re := regexp.MustCompile(`\(up\)[.,!?:;]*`)

	for i := 1; i < len(words); i++ {
		if re.MatchString(words[i]) {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
