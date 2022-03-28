package main

import "errors"

func ProduceExactMatchMatrix(pentaminos []Pentamino, board []BoardCell) EMMatrix {
	emMatrix := EMMatrix{
		make(map[int]int),
		make(map[int]map[int]bool),
	}
	rowIndex := 0
	for pIdx, p := range pentaminos {
		for _, pp := range p.Permutations {
			for _, bc := range board {
				bcs, err := EMCanPlacePentamino(board, pp, bc.Pos)
				if err != nil {
					continue
				}
				row := make(map[int]bool)
				for i := 0; i < 72; i++ {
					if i < 12 {
						if i == pIdx {
							row[i] = true
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
						if bcP.Index+12 == i {
							row[i] = true
							_, ok := emMatrix.ColCount[i]
							if ok {
								emMatrix.ColCount[i] += 1
							} else {
								emMatrix.ColCount[i] = 1
							}
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

type EMCell struct {
	IX   int
	Flag bool
}

type EMMatrix struct {
	ColCount map[int]int
	Matrix   map[int]map[int]bool
}
