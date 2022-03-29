package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main() {
	fmt.Println("Pentomino puzzle 6x10")
	c := make(chan []int)
	wg := new(sync.WaitGroup)
	fmt.Println()
	pentaminos := GeneratePentaminoes()
	RemoveVRotations(pentaminos)
	board := GetBoard6x10()
	emMatrix := ProduceExactMatchMatrix(pentaminos, board)
	matrices := MatricesWithXAnchorPoints(emMatrix, pentaminos[0].Permutations[0])
	fmt.Printf("Em length: %v \n", len(matrices[0].Matrix))

	for _, m := range matrices {
		wg.Add(1)
		go func(matrix EMMatrix) {
			defer wg.Done()
			emStart(matrix, []int{}, c)
		}(m)
	}
	// go func(matrix EMMatrix) {
	// 	defer wg.Done()
	// 	emStart(matrix, []int{}, c)
	// }(emMatrix)
	// for i, b := range boards {
	// 	// PrintBoard(b)
	// 	wg.Add(1)
	// 	go func(board [6][10]string, boardId int) {
	// 		defer wg.Done()
	// 		attempts := make(map[string]int)
	// 		start(pentaminoes, &board, c, boardId, attempts)
	// 	}(b, i)
	// 	// fmt.Println("-----------")
	// }
	go func() {
		wg.Wait()
		close(c)
	}()
	fmt.Println("Started looking for solutions")
	failed := 0
	success := 0
	duplicate := 0
	result := [][]int{}
	fmt.Printf("SolutionsFound: %v. Duplicates Found: %v, Failed attempts: %v.", success, duplicate, failed)

	for rows := range c {
		if len(rows) == 0 {
			failed++
			fmt.Printf("\rSolutionsFound: %v. Duplicates Found: %v, Failed attempts: %v.", success, duplicate, failed)
			continue
		}
		sort.Ints(rows)
		isDuplicate := false
		for _, found := range result {
			if compareResultArr(found, rows) {
				//duplicate
				isDuplicate = true
				break
			}
		}
		if isDuplicate {
			duplicate++
		} else {
			result = append(result, rows)
			success++
		}
		fmt.Printf("\rSolutionsFound: %v. Duplicates Found: %v, Failed attempts: %v.", success, duplicate, failed)
	}
	// for r := range c {
	// 	if r == "failed" {
	// 		failed++
	// 		fmt.Printf("\rSolutionsFound: %v. Failed attempts: %v.", success, failed)
	// 		continue
	// 	}
	// 	result = append(result, r)
	// 	success++
	// 	fmt.Printf("\rSolutionsFound: %v. Failed attempts: %v.", success, failed)
	// }

	fmt.Println()
	fmt.Printf("Number of solutions found: %v\n", success)
}

func compareResultArr(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i, val := range first {
		if val != second[i] {
			return false
		}
	}
	return true
}

func start(pentaminoesLeft []Pentamino, board *[6][10]string, c chan<- string) {
	if len(pentaminoesLeft) == 0 {
		// Board is filled send message on channel
		str2 := BoardToString(*board)
		c <- str2
		return
	}
	for i, p := range pentaminoesLeft {
		for y, row := range board {
			for x, _ := range row {
				for _, pv := range p.Permutations {
					pos := Vector2{x, y}
					if CanPlacePentamino(*board, pv, pos) {
						PlacePentamino(pv, board, p.Id, pos)
						if ValidateBoard(*board) {
							pl1 := make([]Pentamino, len(pentaminoesLeft))
							copy(pl1, pentaminoesLeft)
							pl := append(pl1[:i], pl1[i+1:]...)
							start(pl, board, c)
						}
						RemovePentamino(pv, board, pos)
					}
				}
			}
		}
	}
	c <- "failed"
}
func remove(p []Pentamino, i int) []Pentamino {
	return append(p[:i], p[i+1:]...)
}

func emStart(emMatrix EMMatrix, chosenRows []int, c chan<- []int) {
	if len(emMatrix.Matrix) == 0 {
		// failed
		c <- make([]int, 0, 0)
		return
	}
	if len(emMatrix.Matrix) == 1 {
		for rowI, row := range emMatrix.Matrix {
			for _, c := range row {
				if c == false {
					return
				}
			}
			successRows := []int{}
			copy(successRows, chosenRows)
			successRows = append(successRows, rowI)
			c <- successRows
			// success
			return
		}
	}
	lowestCount := math.MaxInt
	var lowestCol int
	for key, count := range emMatrix.ColCount {
		if count < lowestCount {
			lowestCount = count
			lowestCol = key
		}
	}
	for rowI, row := range emMatrix.Matrix {
		if val, exist := row[lowestCol]; exist && val {
			newMatrix := createNewMatrix(emMatrix, row)
			newChosenRows := []int{}
			newChosenRows = append(newChosenRows, chosenRows...)
			newChosenRows = append(newChosenRows, rowI)
			emStart(newMatrix, newChosenRows, c)

		}
	}
}

func createNewMatrix(matrix EMMatrix, row map[int]bool) EMMatrix {
	colIxToRemove := []int{}
	newMatrix := CopyMatrix(matrix)

	for colI, col := range row {
		if col {
			colIxToRemove = append(colIxToRemove, colI)
		}
	}
	for rowI, rowInner := range matrix.Matrix {
		shouldRemoveRow := false
		for _, colIx := range colIxToRemove {
			if val, hasVal := rowInner[colIx]; hasVal && val {
				shouldRemoveRow = true
				break
			}
		}
		if shouldRemoveRow {
			for colIx, val := range rowInner {
				if val {
					newMatrix.ColCount[colIx] -= 1
					if newMatrix.ColCount[colIx] == 0 {
						delete(newMatrix.ColCount, colIx)
					}
				}
			}
			delete(newMatrix.Matrix, rowI)
		}
	}
	for _, r := range newMatrix.Matrix {
		for _, colIx := range colIxToRemove {
			delete(r, colIx)
		}
	}
	return newMatrix
}
