package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Words struct {
	Words []string `yaml:"words"`
}

var hangmanFigures = []string{
	`
	--------
	 |
	 |
	 |
	 |
	 |
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |
	 |
	 |
	 |
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |
	 |
	 |
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |     |
	 |
	 |
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |    /|
	 |
	 |
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |    /|\
	 |
	 |
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |    /|\
	 |	   |
	 |  
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |    /|\
	 |     |
	 |    / 
	 |
	--------
	 |    |
	`,
	`
	--------
	 |     |
	 |     0
	 |    /|\
	 |     |
	 |    / \
	 |
	--------
	 |    |
	`,
}

func main() {
	words, err := readWordsFromYAML("./hangman/words.yaml")
	if err != nil {
		fmt.Printf("Error reading words from YAML file: %v\n", err)
	}

	if len(words) == 0 {
		log.Fatalf("No words found in YAML file.")
	}

	targetWord := words[rand.Intn(len(words))]
	guessedWord := make([]rune, len(targetWord))
	for i := range guessedWord {
		guessedWord[i] = '_'
	}
	attempts := len(hangmanFigures) - 1
	failedAttempts := 0
	guessedLetters := make(map[rune]bool)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Hangman")
	fmt.Printf("You have %d attempts to guess the word.\n", attempts)

	for failedAttempts < attempts {
		fmt.Printf("Word: %s\n", displayWord(targetWord, guessedLetters))
		fmt.Print("Guess a letter: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		guess := rune(input[0])
		if guessedLetters[guess] {
			fmt.Println("You already guessed that letter.")
		}

		guessedLetters[guess] = true
		correctGuess := false

		for i, char := range targetWord {
			if char == guess {
				guessedWord[i] = guess
				correctGuess = true
			}
		}

		if correctGuess {
			fmt.Println("Good guess? - Good guess!!")
		} else {
			failedAttempts++
			fmt.Println("Nope. Try again.")
			fmt.Println(hangmanFigures[failedAttempts])
			fmt.Printf("You have %d attempts left.\n", attempts-failedAttempts)
		}

		if string(guessedWord) == targetWord {
			fmt.Println("Congratulations!! You have guessed the word:", targetWord)
			return
		}
	}
	fmt.Println(hangmanFigures[failedAttempts])
	fmt.Println("Game over. The word was:", targetWord)
}

func readWordsFromYAML(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var words Words
	err = yaml.Unmarshal(data, &words)
	if err != nil {
		return nil, err
	}

	return words.Words, nil
}

func displayWord(word string, guessedLetters map[rune]bool) string {
	displayedWord := []rune(word)
	for i, char := range displayedWord {
		if !guessedLetters[char] {
			displayedWord[i] = '_'
		}
	}
	return string(displayedWord)
}
