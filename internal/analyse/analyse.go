package analyse

func NumDoubleLetters(words []string) int {
	var wc int
	for _, word := range words {
		if hasDoubleLetters(word) {
			wc++
		}
	}
	return wc
}

func hasDoubleLetters(word string) bool {
	var prevRune rune
	for _, rune := range word {
		if rune == prevRune {
			return true
		}
		prevRune = rune
	}
	return false
}
