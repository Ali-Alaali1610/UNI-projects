package piscine

import (
	"regexp"
	"strconv"
	"strings"
)

func LowerNum(sentence string) string {

	reg := regexp.MustCompile(`\(low,\s*(\d+)\s*\)`)
	location := reg.FindStringIndex(sentence)

	for location != nil {
		word := sentence[location[0]:location[1]]
		numStr := word[6 : len(word)-1]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "Error: Invalid number in (low, <number>)"
		}

		words := strings.Fields(sentence[:location[0]])

		startIndex := len(words) - num
		if startIndex < 0 {
			startIndex = 0
		}

		for i := startIndex; i < len(words); i++ {
			words[i] = strings.ToLower(words[i])
		}

		sentence = strings.Join(words, " ") + sentence[location[1]:]

		location = reg.FindStringIndex(sentence)
	}
	return sentence
}
