package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {
	options := []string{"rock", "paper", "scissors", "lizard", "spock", "rules", "quit"}
	for {
		input := getUserInput()
		if validateInput(options, input) {
			if input == "rules" {
				fmt.Println("Game rules: Rock crushes Scissors,\n" +
					"Scissors cuts Paper, Paper covers Rock,\n" +
					"Rock crushes Lizard, Lizard poisons Spock,\n" +
					"Spock smashes Scissors, Scissors decapitates Lizard,\n" +
					"Lizard eats Paper, Paper disproves Spock, Spock vaporizes Rock.")
			} else if input == "quit" {
				fmt.Println("Exiting the game")
				break
			} else {
				npc := options[rand.Intn(len(options)-2)]
				checkWin(npc, input)
			}

		}
		time.Sleep(10 * time.Millisecond)
	}

}

func getUserInput() string {
	var i string
	fmt.Println("Rock Paper Scissor Lizard Spock! ")
	fmt.Print("Enter \"rules\" or \"quit\" or make your choice to start playing:")
	fmt.Scan(&i)
	return i
}

func validateInput(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

func checkWin(npc string, input string) {
	winConditions := map[string][]string{
		"rock":     {"scissors", "lizard"},
		"paper":    {"rock", "spock"},
		"scissors": {"paper", "lizard"},
		"lizard":   {"spock", "paper"},
		"spock":    {"scissors", "rock"},
	}
	colorInput := green + input + reset
	colorNpc := red + npc + reset
	drawInput := yellow + input + reset
	if npc == input {
		fmt.Println("Draw! you both picked: " + drawInput)
		return
	}

	for _, win := range winConditions[input] {
		if win == npc {
			fmt.Println("You win! " + colorInput + " beats " + colorNpc)
			return
		}
	}
	fmt.Println("You lose! " + colorNpc + " beats " + colorInput)
}
