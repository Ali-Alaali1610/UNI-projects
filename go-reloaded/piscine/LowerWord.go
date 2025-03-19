package piscine

import (
	"regexp"
	"strings"
)

func LowerWord(words []string) []string {
	re := regexp.MustCompile(`\(low\)[.,!?:;]*`)

	for i := 1; i < len(words); i++ {
		if re.MatchString(words[i]) {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
