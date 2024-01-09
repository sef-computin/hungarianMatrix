package main

import (
	"fmt"
	"math/rand"
)

func input(gflag bool) (int, *[][]int) {
	var size int
	var matrix [][]int
	fmt.Print("Введите размер матрицы: ")
	_, _ = fmt.Scanf("%d", &size)
	if gflag {
		matrix = *generateMatrix(size)
	} else {
		fmt.Println("Заполните матрицу")
		matrix = *fillMatrix(size)
	}

	return size, &matrix
}

func generateMatrix(size int) *[][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Int() % 50
		}
	}
	return &matrix
}

func fillMatrix(size int) *[][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}
	return &matrix
}
