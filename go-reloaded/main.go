package main

import (
	"fmt"
	"go-reloaded/piscine"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("the input must be in the form: go run . input.txt output.txt")
		return
	}

	input, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	strInput := string(input)
	words := strings.Fields(strInput)

	words = piscine.HexConversion(words)
	words = piscine.BinConversion(words)
	words = piscine.UpperWord(words)
	words = piscine.LowerWord(words)
	words = piscine.Cap(words)
	words = piscine.BringPunctuation(words)
	words = piscine.CheckForVowels(words)

	sentence := strings.Join(words, " ")

	sentence = piscine.UpperNum(sentence)
	sentence = piscine.LowerNum(sentence)
	sentence = piscine.CapNum(sentence)
	sentence = piscine.TrimApostrophes(sentence)

	file_err := os.WriteFile("result.txt", []byte(sentence), 0644)
	if file_err != nil {
		fmt.Println("error writing in the file", file_err)
	} else {
		fmt.Println("File written successfully!")
	}
}
