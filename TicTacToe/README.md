# Tic-Tac-Toe

This is a simple console game where two players take turns marking spaces in a 3×3 grid, aiming to place three of their marks in a horizontal, vertical, or diagonal row.

## How to Play

1. The game starts with an empty 3×3 grid.
2. Two players take turns to place their marks (X and O) in an empty cell.
3. The first player to get three of their marks in a row (horizontally, vertically, or diagonally) wins the game.
4. If all nine cells are filled and neither player has three marks in a row, the game ends in a draw.

## Running the Game

To run the game, navigate to the `TicTacToe` directory, download the dependencies, and execute the program:

```sh
go mod tidy
go run .