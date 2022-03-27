package main

import "errors"

func ProduceExactMatchMatrix(pentaminos []Pentamino, board []BoardCell) [][72]bool {
	emMatrix := [][72]bool{}
	for pIdx, p := range pentaminos {
		for _, pp := range p.Permutations {
			for _, bc := range board {
				bcs, err := EMCanPlacePentamino(board, pp, bc.Pos)
				if err != nil {
					continue
				}
				row := [72]bool{}
				for i := range row {
					if i < 12 {
						row[i] = i == pIdx
						continue
					}
					for _, bcP := range bcs {
						if bcP.Index == i+12 {
							row[i] = true
						}
					}
				}
				emMatrix = append(emMatrix, row)
			}
		}
	}
	return emMatrix
}

type BoardCell struct {
	Pos   Vector2
	Index int
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
