package main

import "fmt"

type selection struct {
	i int
	j int
}
type point struct {
	x int
	y int
}

func solve(matrix [][]int, max_task bool) [][]int {
	if max_task {
		minimizationTask(matrix)
	}
	siz, k := initSIZ(matrix)

	n := len(matrix)

	sel := []selection{}

	sel = initSelections(matrix, sel, siz)

	// count := 0
	for k < n {
		// count++
		fmt.Println("k=", k, "n=", n)

		fmt.Println("Матрица:")
		show(matrix, n)
		fmt.Println("Текущая СНН:")
		showSIZ(siz, n)
		fmt.Print("Выделения: ")
		fmt.Println(sel)

		if hasUnselectedZeroes(matrix, sel) {
			fmt.Println("Есть невыделенные нули")
			k = processUnselectedZeroes(matrix, &sel, siz, k)
		} else {
			fmt.Println("Нет невыделенных нулей")
			processNoUnselectedZeroes(matrix, sel)
		}
	}

	return siz
}

func processNoUnselectedZeroes(matrix [][]int, sel []selection) {
	h := -1
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if !rowSelected(sel, i) && !colSelected(sel, j) {
				if h == -1 || matrix[i][j] < h {
					h = matrix[i][j]
				}
			}
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if !colSelected(sel, j) {
				matrix[i][j] -= h
			}
			if rowSelected(sel, i) {
				matrix[i][j] += h
			}
		}
	}

}

func minimizationTask(matrix [][]int) {
	size := len(matrix)
	max := matrix[0][0]
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
	}

	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			matrix[i][j] = max - matrix[i][j]
		}
	}
}

func prepare(matrix [][]int) {
	prepareCols(matrix)
	prepareRows(matrix)
}

func prepareCols(matrix [][]int) {
	size := len(matrix)
	for j := 0; j < size; j++ {
		min := matrix[0][j]
		for i := 0; i < size; i++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}

		for i := 0; i < size; i++ {
			matrix[i][j] -= min
		}
	}
}

func prepareRows(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		min := matrix[i][0]
		for j := 0; j < size; j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}

		for j := 0; j < size; j++ {
			matrix[i][j] -= min
		}
	}
}

func initSelections(matrix [][]int, sel []selection, siz [][]int) []selection {
	sel = make([]selection, 0)
	size := len(matrix)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if siz[i][j] == 1 {
				sel = append(sel, selection{i: -1, j: j})
			}
		}
	}
	return sel
}

func hasUnselectedZeroes(matrix [][]int, sel []selection) bool {
	size := len(matrix)
	for j := 0; j < size; j++ {
		if colSelected(sel, j) {
			continue
		}
		for i := 0; i < size; i++ {
			if rowSelected(sel, i) {
				continue
			}
			if matrix[i][j] == 0 {
				fmt.Println("Невыделенный 0: ", i, j)
				return true
			}
		}
	}
	return false
}

func processUnselectedZeroes(matrix [][]int, sel *[]selection, siz [][]int, k int) int {
	size := len(matrix)

	for j := 0; j < size; j++ {
		if colSelected(*sel, j) {
			continue
		}
		for i := 0; i < size; i++ {
			if rowSelected(*sel, i) {
				continue
			}
			if matrix[i][j] == 0 {
				siz[i][j] = 2
				showSIZ(siz, size)
				if col := getZeroStar(i, siz); col != -1 {
					*sel = deleteColSelection(*sel, col)
					*sel = setRowSelection(*sel, i)

					fmt.Print("Выделения: ")
					fmt.Println(*sel)
				} else {
					fmt.Println("Строим L-цепочку c позиции", i, j)
					lChain := make([][2]int, 0)
					createLChain(&lChain, siz, map[point]bool{}, i, j, 1)
					fmt.Println("L-цепочка построена")
					fmt.Println(lChain)
					for _, val := range lChain {
						siz[val[0]][val[1]]--
					}
					*sel = initSelections(matrix, *sel, siz)

					k++
					return k
				}
			}
		}
	}
	return k
}

func createLChain(lChain *[][2]int, siz [][]int, visited map[point]bool, i, j int, current_wanted_elem int) {
	*lChain = append(*lChain, [2]int{i, j})
	if x, y, ok := isPresent(siz, i, j, current_wanted_elem); ok {
		p := point{x, y}
		if !visited[p] {

			visited[p] = true
			current_wanted_elem *= -1
			createLChain(lChain, siz, visited, x, y, current_wanted_elem)
		}

	} else {
		return
	}
}

func isPresent(siz [][]int, i, j, current_wanted_elem int) (int, int, bool) {
	size := len(siz)
	switch current_wanted_elem {
	case 1:
		for k := 0; k < size; k++ {
			if siz[k][j] == 1 {
				fmt.Println("0* найден в", k, j)
				return k, j, true
			}
		}
	case -1:
		for k := 0; k < size; k++ {
			if siz[i][k] == 2 {
				fmt.Println("0' найден в", i, k)
				return i, k, true
			}
		}
	}
	return -1, -1, false
}

func getZeroStar(row int, siz [][]int) int {
	for j := 0; j < len(siz[row]); j++ {
		if siz[row][j] == 1 {
			return j
		}
	}
	return -1

}

func deleteColSelection(sel []selection, col int) []selection {
	ind := -1
	for i, val := range sel {
		if val.j == col {
			ind = i
		}
	}

	if ind != -1 {
		sel = append(sel[:ind], sel[ind+1:]...)
	}
	return sel
}

func setRowSelection(sel []selection, row int) []selection {
	sel = append(sel, selection{i: row, j: -1})
	return sel
}

func colSelected(sel []selection, j int) bool {
	for _, val := range sel {
		if val.j == j {
			return true
		}
	}
	return false
}

func rowSelected(sel []selection, i int) bool {
	for _, val := range sel {
		if val.i == i {
			return true
		}
	}
	return false
}

func initSIZ(matrix [][]int) ([][]int, int) {
	size := len(matrix)
	siz := make([][]int, size)
	for i := 0; i < size; i++ {
		siz[i] = make([]int, size)
	}
	var k int
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if matrix[i][j] == 0 {
				if func() bool {
					for k := 0; k < size; k++ {
						if siz[k][j] == 1 || siz[i][k] == 1 {
							return false
						}
					}
					return true
				}() {
					siz[i][j] = 1
					k++
				}
			}
		}
	}

	return siz, k
}
