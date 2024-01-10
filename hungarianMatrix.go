package main

import "fmt"

func main() {
	_, matrixp := input(true)
	matrix := *matrixp
	fmt.Println("Изначальная матрица стоимостей")
	showMatrix(matrix)

	prepare(matrix)
	fmt.Println("Приведенная матрица:")
	showMatrix(matrix)

	ans := solve(matrix, false)
	fmt.Println("Финальная СНН:")
	showAnswer(ans)
}
