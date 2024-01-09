package main

import "fmt"

func show(matrix [][]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print("|")
		for j := 0; j < size; j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Print(" |\n")
	}
}

func showSIZ(siz [][]int, size int) {
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

func showAnswer(siz [][]int, size int) {
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
