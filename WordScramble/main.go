package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	reset  = "\033[0m"
	orange = "\033[38;5;208m"
)

func main() {
	var i string
	game := 1
	for game == 1 {
		word := getWord()
		scrambledWord := scrambleWord(word)
		for {
			fmt.Println("Guess the word: " + orange + scrambledWord + reset)
			fmt.Scan(&i)
			if i == "quit" {
				fmt.Println("Exiting game")
				game = 0
				break
			} else if i == word {
				fmt.Println("You won!")
				break
			} else {
				fmt.Println("No sorry! Try again.")
			}
		}
	}
}

func getWord() string {
	words := []string{"example", "word", "scramble", "game", "play"}
	return words[rand.Intn(len(words))]
}

func scrambleWord(word string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	scrambled := []rune(word)
	r.Shuffle(len(scrambled), func(i, j int) { scrambled[i], scrambled[j] = scrambled[j], scrambled[i] })
	return string(scrambled)
}
