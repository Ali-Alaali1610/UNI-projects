package piscine

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CapNum(sentence string) string {

	reg := regexp.MustCompile(`\(cap,\s*(\d+)\s*\)`)
	location := reg.FindStringIndex(sentence)

	for location != nil {
		word := sentence[location[0]:location[1]]
		numStr := word[6 : len(word)-1]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "Error: Invalid number in (cap, <number>)"
		}

		words := strings.Fields(sentence[:location[0]])

		startIndex := len(words) - num
		if startIndex < 0 {
			startIndex = 0
		}

		for i := startIndex; i < len(words); i++ {

			if len(words[i]) > 0 && words[i][0] >= 'a' && words[i][0] <= 'z' {
				words[i] = strings.ToUpper(string(words[i][0])) + words[i][1:]
			} else{
				fmt.Println("cant capitalize:", string(words[i][0]))
			}
		}

		sentence = strings.Join(words, " ") + sentence[location[1]:]

		location = reg.FindStringIndex(sentence)
	}
	return sentence
}
