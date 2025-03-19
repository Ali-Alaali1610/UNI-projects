package piscine

import (
	"fmt"
	"regexp"
	"strconv"
)

func BinConversion(words []string) []string {
	re := regexp.MustCompile(`\(bin\)[.,!?:;]*`)

	for i := 1; i < len(words); i++ {
		if re.MatchString(words[i]) {

			decimal, err := strconv.ParseInt(words[i-1], 2, 64)

			if err != nil {
				fmt.Println("Error:", err)
				return words
			}

			words[i-1] = strconv.Itoa(int(decimal))
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
