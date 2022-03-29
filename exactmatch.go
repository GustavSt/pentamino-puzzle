package main

import "errors"

func ProduceExactMatchMatrix(pentaminos []Pentamino, board []BoardCell) EMMatrix {
	emMatrix := EMMatrix{
		make(map[int]int),
		make(map[int]map[int]bool),
	}
	rowIndex := 0
	addedBoardCells := [][5]BoardCell{}
	for pIdx, p := range pentaminos {
		for _, pp := range p.Permutations {
			for _, bc := range board {
				bcs, err := EMCanPlacePentamino(board, pp, bc.Pos)
				if err != nil {
					continue
				}
				addedBoardCells = append(addedBoardCells, bcs)
				row := make(map[int]bool)
				for i := 0; i < 72; i++ {
					if i < 12 {
						row[i] = i == pIdx
						if row[i] {
							_, ok := emMatrix.ColCount[i]
							if ok {
								emMatrix.ColCount[i] += 1
							} else {
								emMatrix.ColCount[i] = 1
							}
						}
						continue
					}
					for _, bcP := range bcs {
						row[i] = bcP.Index+12 == i
						if row[i] {
							_, ok := emMatrix.ColCount[i]
							if ok {
								emMatrix.ColCount[i] += 1
							} else {
								emMatrix.ColCount[i] = 1
							}
							break
						}
					}
				}
				emMatrix.Matrix[rowIndex] = row
				rowIndex++
			}
		}
	}
	return emMatrix
}

func MatricesWithXAnchorPoints(matrix EMMatrix, p [5]Vector2) []EMMatrix {
	anchorPoints := GetXAncorPoints()
	result := []EMMatrix{}
	for _, ap := range anchorPoints {
		newMatrix := CopyMatrix(matrix)
		for rowI, row := range matrix.Matrix {
			if row[0] { // 0 is X pentamino
				delete(newMatrix.Matrix, rowI)
				for ci, c := range row {
					if c {
						newMatrix.ColCount[ci] -= 1
					}
				}
			}
		}
		row := getMatrixRow(ap, p)
		for colIx, c := range row {
			if c {
				_, ok := newMatrix.ColCount[colIx]
				if ok {
					newMatrix.ColCount[colIx] += 1
				} else {
					newMatrix.ColCount[colIx] = 1
				}
			}
		}
		newMatrix.Matrix[0] = row
		result = append(result, newMatrix)
	}
	return result
}

func getMatrixRow(pos Vector2, pent [5]Vector2) map[int]bool {
	row := make(map[int]bool)
	for i := 0; i < 72; i++ {
		if i < 12 {
			// Make 0(x pentamino) true, others false
			row[i] = i == 0
			continue
		}
		for _, p := range pent {
			x := pos.X + p.X
			y := pos.Y + p.Y
			ix := y*10 + x

			if ix+12 == i {
				row[i] = true
				break
			}
		}
	}
	return row
}

func SelectPos(board []BoardCell) []Vector2 {
	res := make([]Vector2, len(board))
	for _, p := range board {
		res = append(res, p.Pos)
	}
	return res
}

func FindPos(board []BoardCell, pos Vector2) (BoardCell, error) {
	for _, bc := range board {
		if bc.Pos.X == pos.X && bc.Pos.Y == pos.Y {
			return bc, nil
		}
	}
	return BoardCell{}, errors.New("Pos not found")
}

func CopyMatrix(matrix EMMatrix) EMMatrix {
	newMatrix := EMMatrix{
		make(map[int]int),
		make(map[int]map[int]bool),
	}
	for key, val := range matrix.ColCount {
		newMatrix.ColCount[key] = val
	}
	for rowI, row := range matrix.Matrix {
		newMatrix.Matrix[rowI] = make(map[int]bool)
		for colI, col := range row {
			newMatrix.Matrix[rowI][colI] = col
		}
	}
	return newMatrix
}

func EMValidateBoard(board []BoardCell, p [5]BoardCell) bool {
	checked := make([]Vector2, 0)
	for _, cell := range board {
		if contains(checked, cell.Pos) {
			continue
		}
		nrOfEmptySpaces := emNrEmptySpaces(cell, board, p, &checked)
		if len(nrOfEmptySpaces)%5 == 0 {
			continue
		}
		return false
	}

	return true
}
func emNrEmptySpaces(cell BoardCell, board []BoardCell, p [5]BoardCell, checked *[]Vector2) []BoardCell {
	if contains(*checked, cell.Pos) {
		return make([]BoardCell, 0)
	}
	currSpaces := []BoardCell{cell}
	*checked = append(*checked, cell.Pos)
	for _, pc := range p {
		if pc.Index == cell.Index {
			return make([]BoardCell, 0)
		}
	}
	y := cell.Index / 10
	x := cell.Index % 10
	if x+1 < 10 {
		newCell := BoardCell{Pos: Vector2{x + 1, y}, Index: cell.Index + 1}
		currSpaces = append(currSpaces, emNrEmptySpaces(newCell, board, p, checked)...)
	}
	if y+1 < 6 {
		newCell := BoardCell{Pos: Vector2{x, y + 1}, Index: cell.Index + 10}
		currSpaces = append(currSpaces, emNrEmptySpaces(newCell, board, p, checked)...)
	}
	if x-1 > -1 {
		newCell := BoardCell{Pos: Vector2{x - 1, y}, Index: cell.Index - 1}
		currSpaces = append(currSpaces, emNrEmptySpaces(newCell, board, p, checked)...)
	}
	if y-1 > -1 {
		newCell := BoardCell{Pos: Vector2{x, y - 1}, Index: cell.Index - 10}
		currSpaces = append(currSpaces, emNrEmptySpaces(newCell, board, p, checked)...)
	}
	return currSpaces
}

func EMCanPlacePentamino(board []BoardCell, p [5]Vector2, pos Vector2) ([5]BoardCell, error) {
	boardPos := SelectPos(board)
	result := [5]BoardCell{}
	for i, pv := range p {
		bPos := Vector2{pos.X + pv.X, pos.Y + pv.Y}
		if contains(boardPos, bPos) {

			bc, err := FindPos(board, bPos)
			if err != nil {
				return result, err
			}
			result[i] = bc
			continue
		} else {
			return result, errors.New("Pos not in board")
		}
	}
	if !EMValidateBoard(board, result) {
		return result, errors.New("Invalid Pos in board")
	}

	return result, nil
}

type EMMatrix struct {
	ColCount map[int]int
	Matrix   map[int]map[int]bool
}
