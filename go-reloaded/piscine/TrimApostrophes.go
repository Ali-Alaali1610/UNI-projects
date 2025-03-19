package piscine

func TrimApostrophes(sentence string) string {

	for i := 0; i < len(sentence); i++ {

		if sentence[i] == '\'' {
			if i+1 < len(sentence) && sentence[i+1] == ' ' {
				sentence = sentence[:i+1] + sentence[i+2:]
			}
			for j := i + 1; j < len(sentence); j++ {

				if sentence[j] == '\'' {
					if sentence[j-1] == ' ' {
						sentence = sentence[:j-1] + sentence[j:]
					}
					i = j
					break
				}
			}
		}
	}
	return sentence
}
