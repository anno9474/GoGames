package main

import (
	"fmt"
	"os"
)

func main() {
	board := [4][4]string{
		{" ", "1", "2", "3"},
		{"a", "-", "-", "-"},
		{"b", "-", "-", "-"},
		{"c", "-", "-", "-"},
	}
	end := 0
	turn := 1
	var player string

	for end == 0 {
		printBoard(board)
		if turn%2 == 0 {
			player = "X"
			fmt.Println("Player two's turn!")
		} else {
			player = "O"
			fmt.Println("Player one's turn!")
		}
		fmt.Print("Enter your move (e.g., 1a, 3b or q to quit): ")

		board = playerInput(board, player)
		end = endGame(board)
		if end != 0 {
			fmt.Println("You won!")
			printBoard(board)

		}
		turn++
	}
}

func printBoard(board [4][4]string) {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Printf("| %s |", board[row][col])
		}
		fmt.Println()
	}
}

func playerInput(board [4][4]string, player string) [4][4]string {
	var input string
	for {
		fmt.Scan(&input)
		if len(input) == 2 {
			row := input[1]
			col := input[0]

			rowIndex, colIndex := -1, -1

			switch row {
			case 'a':
				rowIndex = 1
			case 'b':
				rowIndex = 2
			case 'c':
				rowIndex = 3
			}

			if col >= '1' && col <= '3' {
				colIndex = int(col - '0')
			}

			if rowIndex != -1 && colIndex != -1 && board[rowIndex][colIndex] == "-" {
				board[rowIndex][colIndex] = player
				break
			}
		} else if input == "q" {
			fmt.Println("Exiting the game.")
			os.Exit(0)
		}

		fmt.Println("Invalid input. Please try again.")
	}
	return board
}

func endGame(board [4][4]string) int {

	for i := 1; i <= 3; i++ {
		if board[i][1] == board[i][2] && board[i][2] == board[i][3] && board[i][1] != "-" {
			if board[i][1] == "X" {
				return 1
			} else {
				return 2
			}
		}
		if board[1][i] == board[2][i] && board[2][i] == board[3][i] && board[1][i] != "-" {
			if board[1][i] == "X" {
				return 1
			} else {
				return 2
			}
		}
	}

	if board[1][1] == board[2][2] && board[2][2] == board[3][3] && board[1][1] != "-" {
		if board[1][1] == "X" {
			return 1
		} else {
			return 2
		}
	}
	if board[1][3] == board[2][2] && board[2][2] == board[3][1] && board[1][3] != "-" {
		if board[1][3] == "X" {
			return 1
		} else {
			return 2
		}
	}

	draw := true
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if board[i][j] == "-" {
				draw = false
				break
			}
		}
	}

	if draw {
		return 3
	}

	return 0
}
