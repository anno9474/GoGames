# Rock Paper Scissors Lizard Spock

This is a simple console game where the player competes against the computer in an expanded version of the classic Rock Paper Scissors game, which includes Lizard and Spock.

## How to Play

1. The player selects either Rock, Paper, Scissors, Lizard, or Spock.
2. The computer randomly selects Rock, Paper, Scissors, Lizard, or Spock.
3. The game determines the winner based on the following rules:
    - Rock crushes Scissors
    - Scissors cuts Paper
    - Paper covers Rock
    - Rock crushes Lizard
    - Lizard poisons Spock
    - Spock smashes Scissors
    - Scissors decapitates Lizard
    - Lizard eats Paper
    - Paper disproves Spock
    - Spock vaporizes Rock
4. The game announces the winner of each round.

## Running the Game

To run the game, navigate to the `RockPaperScissorsLizardSpock` directory, download the dependencies, and execute the program:

```sh
go mod tidy
go run .