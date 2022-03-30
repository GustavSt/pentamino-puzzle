package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main() {
	fmt.Println("Pentomino puzzle 6x10")
	c := make(chan ECMessage)
	wg := new(sync.WaitGroup)
	pentaminos := GeneratePentaminoes()
	RemoveVRotations(pentaminos)
	board := GetBoard6x10()
	ecMatrix := ProduceExactCoverMatrix(pentaminos, board)
	matrices := MatricesWithXAnchorPoints(ecMatrix, pentaminos[0].Permutations[0])
	fmt.Printf("Em length: %v \n", len(matrices[0].Matrix))

	failedCount := [7]int{}
	successCount := [7]int{}
	duplicateCount := [7]int{}

	for i, m := range matrices {
		wg.Add(1)
		failedCount[i] = 0
		successCount[i] = 0
		duplicateCount[i] = 0
		go func(goroutineId int, matrix ECMatrix) {
			defer wg.Done()
			ecStart(matrix, []int{}, c, goroutineId)
		}(i, m)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	fmt.Println("Started looking for solutions")
	failed := 0
	success := 0
	duplicate := 0
	result := []ECMessage{}
	fmt.Printf("SolutionsFound: %v. Duplicates Found: %v, Failed attempts: %v.", success, duplicate, failed)

	for message := range c {
		if len(message.Rows) == 0 {
			failed++
			failedCount[message.Id] += 1
			// printSeparateGoRoutines(successCount, duplicateCount, failedCount)
			fmt.Printf("\rSolutionsFound: %v. Duplicates Found: %v, Failed attempts: %v.", success, duplicate, failed)
			continue
		}
		sort.Ints(message.Rows)
		isDuplicate := false
		for _, found := range result {
			if compareResultArr(found.Rows, message.Rows) {
				isDuplicate = true
				break
			}
		}
		if isDuplicate {
			duplicate++
			duplicateCount[message.Id] += 1
		} else {
			result = append(result, message)
			successCount[message.Id] += 1
			success++
		}
		// printSeparateGoRoutines(successCount, duplicateCount, failedCount)
		fmt.Printf("\rSolutionsFound: %v. Duplicates Found: %v, Failed attempts: %v.", success, duplicate, failed)
	}
	fmt.Println()
	fmt.Printf("Number of solutions found: %v\n", len(result))
}

func printSeparateGoRoutines(successCount [7]int, duplicateCount [7]int, failedCount [7]int) {
	str := fmt.Sprintf("S: ")
	for _, val := range successCount {
		str += fmt.Sprintf("c: %v|", val)
	}
	str += fmt.Sprintf("D: ")
	for _, val := range duplicateCount {
		str += fmt.Sprintf("c: %v|", val)
	}
	str += fmt.Sprintf("F: ")
	for _, val := range failedCount {
		str += fmt.Sprintf("c: %v|", val)
	}
	fmt.Printf("\r %v", str)
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

// Recursive function testing every pentamino on every cell
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

type ECMessage struct {
	Id   int
	Rows []int
}

// recursive function using exact cover algorithm, supposedly.
func ecStart(emMatrix ECMatrix, chosenRows []int, c chan<- ECMessage, goRoutineId int) {
	if len(emMatrix.Matrix) == 0 {
		// failed
		c <- ECMessage{goRoutineId, make([]int, 0, 0)}
		return
	}
	if len(emMatrix.Matrix) == 1 {
		for rowI, row := range emMatrix.Matrix {
			for _, c := range row {
				if c == false {
					// failed
					return
				}
			}
			successRows := []int{}
			successRows = append(successRows, chosenRows...)
			successRows = append(successRows, rowI)
			c <- ECMessage{goRoutineId, successRows}
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
			ecStart(newMatrix, newChosenRows, c, goRoutineId)
		}
	}
}

func createNewMatrix(matrix ECMatrix, row map[int]bool) ECMatrix {
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
