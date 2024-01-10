package main

import "fmt"

func showMatrix(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		fmt.Print("|")
		for j := 0; j < size; j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Print(" |\n")
	}
}

func showSelections(matrix [][]int, sel []selection) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		if rowSelected(sel, i) {
			fmt.Print("+")
		} else {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for j := 0; j < size; j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Print(" |\n")
	}
	fmt.Print("  ")
	for j := 0; j < size; j++ {
		if colSelected(sel, j) {
			fmt.Printf("%4s", "+")
		}
	}
	fmt.Println()
}

func showSIZ(siz [][]int) {
	size := len(siz)
	for i := 0; i < size; i++ {
		fmt.Print("|")
		for j := 0; j < size; j++ {
			if siz[i][j] == 1 {
				fmt.Printf("%4s", "0*")
			} else if siz[i][j] == 2 {
				fmt.Printf("%4s", "0'")
			} else {
				fmt.Printf("%4d", 0)
			}

		}
		fmt.Print(" |\n")
	}
}

func showAnswer(siz [][]int) {
	size := len(siz)
	for i := 0; i < size; i++ {
		fmt.Print("|")
		for j := 0; j < size; j++ {
			if siz[i][j] == 1 {
				fmt.Printf("%4s", "0*")
			} else {
				fmt.Printf("%4d", 0)
			}
		}
		fmt.Print(" |\n")
	}
}
