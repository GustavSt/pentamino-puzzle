package main

import "fmt"

func main() {
	fmt.Println("Pentomino puzzle 6x10")
	boards, _ := GetAnchorBoards()
	for _, b := range boards {
		PrintBoard(b)
		fmt.Println("-----------")
	}
}

func start(pentaminoesLeft []Pentamino, board *[6][10]string) {
	if len(pentaminoesLeft) == 0 {
		// Board is filled send message on channel
		return
	}
	for i, p := range pentaminoesLeft {
		for _, pv := range p.Permutations {
			for y, row := range board {
				for x, _ := range row {
					pos := Vector2{x, y}
					if CanPlacePentamino(board, pv, pos) {
						PlacePentamino(pv, board, p.Id, pos)
						if ValidateBoard(board) {
							pl := append(pentaminoesLeft[:i], pentaminoesLeft[i+1:]...)
							start(pl, board)
						}
						RemovePentamino(pv, board, pos)
					}
				}
			}
		}
	}
}
