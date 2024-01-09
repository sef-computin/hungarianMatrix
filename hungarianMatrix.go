package main

import "fmt"

func main() {
	size, matrixp := input(false)
	matrix := *matrixp
	fmt.Println("Изначальная матрица стоимостей")
	show(matrix, size)

	prepare(matrix)
	fmt.Println("Приведенная матрица:")
	show(matrix, size)

	ans := solve(matrix, false)
	fmt.Println("Финальная СНН:")
	showAnswer(ans, size)
}
