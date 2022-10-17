package words

import (
	"strings"
)

func GetWords() ([]string, []string) {
	words := string(allWords)
	words = strings.ReplaceAll(words, "\"", "")
	wordlist := strings.Split(words, ",")

	wordlelist := strings.Fields(string(wordles))

	return wordlist, wordlelist
}
