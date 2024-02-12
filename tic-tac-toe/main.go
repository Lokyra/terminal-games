package main

import (
	"fmt"
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

func main() {
	fmt.Println("Hello World!")

	//	var numbers [10]int
	// var matrix [3][3]int
}
