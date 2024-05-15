package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var maxRandomNumber = 100
	var minRandomNumber = 1
	var input int

	randomNumber := rand.IntN(maxRandomNumber-minRandomNumber) + minRandomNumber

	for {
		fmt.Print("Guess the number: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			_, _ = fmt.Scanln()
			continue
		}

		if input <= 0 || input >= 100 {
			fmt.Println("Please enter a valid number from 1 to 100")
		} else if input > randomNumber {
			fmt.Println("Your guess is too high!")
		} else if input < randomNumber {
			fmt.Println("Your guess is too low!")
		} else if input == randomNumber {
			fmt.Println("Your guess is correct!")
			break
		}
	}
}
