package piscine

import (
	"regexp"
)

func BringPunctuation(words []string) []string {
	re := regexp.MustCompile(`^[.,!?:;]+$`)

	for i := 0; i < len(words); i++ {
		if re.MatchString(words[i]) {
			word := words[i]
			if i > 0 {
				words[i-1] += word
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	for i := 1; i < len(words); i++ {
		j := 0
		for j < len(words[i]) && (words[i][j] == '.' || words[i][j] == ',' || words[i][j] == '!' ||
			words[i][j] == '?' || words[i][j] == ':' || words[i][j] == ';') {

			words[i-1] += string(words[i][j])

			words[i] = words[i][:j] + words[i][j+1:]

		}
	}
	return words
}
