# Hangman

This is a simple console game where the player tries to guess a randomly picked word from a list.

## How to Play

1. The computer selects a random word from a list.
2. The player has to guess the word by suggesting letters one at a time.
3. The computer provides feedback indicating which letters are correct and their positions in the word.
4. The game continues until the player guesses the correct word or reaches the maximum number of incorrect guesses.

## Running the Game

To run the game, navigate to the `Hangman` directory, download the dependencies, and execute the program:

```sh
go mod tidy
go run .
