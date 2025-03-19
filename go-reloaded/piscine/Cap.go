package piscine

import (
	"fmt"
	"regexp"
	"strings"
)

func Cap(words []string) []string {

	re := regexp.MustCompile(`\(cap\)[.,!?:;]*`)

	for i := 1; i < len(words); i++ {
		if re.MatchString(words[i]) {

			if words[i-1][0] >= 'a' && words[i-1][0] <= 'z' {
				words[i-1] = strings.ToUpper(string(words[i-1][0])) + words[i-1][1:]
			} else{
				fmt.Println("cant capitalize:", string(words[i-1][0]))
			}
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
