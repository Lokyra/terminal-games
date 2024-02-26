package main

import (
	"fmt"
	"math/rand"
)

func welcome() {
	fmt.Println("Welcome To Tic-Tac-Toe games")
}

func printArray(numbers [10]int) {
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
		fmt.Println(numbers[i])
	}
}

func printMatix(matrix [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = j + 1
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func getPosition() (int, int) {
	var x, y int
	fmt.Print("\nRow : ")
	fmt.Scan(&x)
	fmt.Print("\nColumn : ")
	fmt.Scan(&y)
	return x, y
}

func getAiPosition(board [3][3]string) (int, int) {
	var x, y int

	for {
		x = rand.Intn(3)
		y = rand.Intn(3)
		if board[x][y] == " " {
			return x, y
		}
	}
}

func changeBoard(board [3][3]string, x int, y int, n int) [3][3]string {
	if n == 1 {
		board[x][y] = "X"
	} else {
		board[x][y] = "O"
	}
	return board
}

func displayBoard(board [3][3]string) {
	fmt.Printf(" %s  | %s  | %s \n", board[0][0], board[0][1], board[0][2])
	fmt.Println("--------------")
	fmt.Printf(" %s  | %s  | %s \n", board[1][0], board[1][1], board[1][2])
	fmt.Println("--------------")
	fmt.Printf(" %s  | %s  | %s ", board[2][0], board[2][1], board[2][2])
}

func horizontalWin(board [3][3]string) bool {
	if board[0][0] == "X" && board[0][0] == board[0][1] && board[0][1] == board[0][2] {
		fmt.Println("\nHuman Won !")
		return false
	}
	if board[1][0] == "X" && board[1][0] == board[1][1] && board[1][1] == board[1][2] {
		fmt.Println("\nHuman Won !")
		return false
	}
	if board[2][0] == "X" && board[2][0] == board[2][1] && board[2][1] == board[2][2] {
		fmt.Println("\nHuman Won !")
		return false
	}
	return true
}

func verticalWin(board [3][3]string) bool {
	if board[0][0] == "X" && board[0][0] == board[1][0] && board[1][0] == board[2][0] {
		fmt.Println("\nHuman Won !")
		return false
	}
	if board[0][1] == "X" && board[0][1] == board[1][1] && board[1][1] == board[2][1] {
		fmt.Println("\nHuman Won !")
		return false
	}
	if board[0][2] == "X" && board[0][2] == board[1][2] && board[1][2] == board[2][2] {
		fmt.Println("\nHuman Won !")
		return false
	}
	return true
}

func diagonalWin(board [3][3]string) bool {
	if board[0][0] == "X" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		fmt.Println("\nHuman Won !")
		return false
	}
	if board[0][2] == "X" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		fmt.Println("\nHuman Won !")
		return false
	}
	return true
}

/*
func checkHumanWin(board [3][3]string) bool {
	run := true
	run = horizontalWin(board)
	if run == false {
		return run
	}
	run = diagonalWin(board)
	if run == false {
		return run
	}
	run = diagonalWin(board)
	if run == false {
		return run
	}
	return run
}
*/

func main() {
	fmt.Println("Welcome to The Tic-Tac-Toe game !")
	fmt.Println()
	var x, y int
	var run bool = true
	board := [3][3]string{
		{" ", " ", " "},

		{" ", " ", " "},

		{" ", " ", " "},
	}

	for {

		fmt.Println("\nHuman turn : ")
		displayBoard(board)
		x, y = getPosition()
		board = changeBoard(board, x, y, 1)
		displayBoard(board)

		fmt.Println()
		fmt.Println()

		fmt.Println("\nMachine turn : ")
		x, y = getAiPosition(board)
		board = changeBoard(board, x, y, 0)
		displayBoard(board)

		fmt.Println()
		fmt.Println()
		run = verticalWin(board)
		if run == false {
			return
		}
		run = horizontalWin(board)
		if run == false {
			return
		}
		run = diagonalWin(board)
		if run == false {
			return
		}
	}
}
