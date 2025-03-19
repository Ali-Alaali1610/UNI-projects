package piscine

func CheckForVowels(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && (words[i+1][0] == 'a' || words[i+1][0] == 'A' ||
			words[i+1][0] == 'e' || words[i][0] == 'E' ||
			words[i+1][0] == 'i' || words[i+1][0] == 'I' ||
			words[i+1][0] == 'o' || words[i+1][0] == 'O' ||
			words[i+1][0] == 'u' || words[i+1][0] == 'U') {

			words[i] += "n"

		}
	}
	return words
}
