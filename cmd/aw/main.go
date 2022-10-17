package main

import (
	"fmt"

	"github.com/psvehla/wordle/internal/analyse"
	"github.com/psvehla/wordle/internal/words"
)

func main() {
	wordlist, wordlelist := words.GetWords()
	fmt.Println("Corpus of all 5 letter words:", len(wordlist))
	fmt.Println("Percentage of double letter words:", analyse.NumDoubleLetters(wordlist)*100/len(wordlist))
	fmt.Println("Corpus of words used by Wordle:", len(wordlelist))
	fmt.Println("Percentage of double letter words:", analyse.NumDoubleLetters(wordlelist)*100/len(wordlelist))
}
