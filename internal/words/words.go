package words

import (
	"os"
	"strings"
)

func GetWords() ([]string, []string) {
	allWords, err := os.ReadFile("assets/word-list.csv")
	words := string(allWords)
	words = strings.ReplaceAll(words, "\"", "")
	if err != nil {
		panic(err)
	}
	wordlist := strings.Split(words, ",")

	wordles, err := os.ReadFile("assets/wordle-list.txt")
	if err != nil {
		panic(err)
	}
	wordlelist := strings.Fields(string(wordles))

	return wordlist, wordlelist
}
