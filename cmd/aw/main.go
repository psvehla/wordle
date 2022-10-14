package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words, err := os.ReadFile("assets/word-list.csv")
	if err != nil {
		panic(err)
	}
	wordlist := strings.Split(string(words), ",")
	fmt.Println(wordlist)
	fmt.Println(len(wordlist))

	wordles, err := os.ReadFile("assets/wordle-list.txt")
	if err != nil {
		panic(err)
	}
	wordlelist := strings.Fields(string(wordles))
	fmt.Println(wordlelist)
	fmt.Println(len(wordlelist))
}
