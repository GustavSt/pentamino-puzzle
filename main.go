package main

import "fmt"

func main() {
	fmt.Println("Pentomino puzzle 6x10")
	board := [6][10]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	penatminoes := generatePentaminoes()
	for _, p := range penatminoes {
		for _, pp := range p.Permutations {
			// Find x,y to place pp
		}
	}
	printBoard(board)
}

func start(pentaminoesLeft []Pentamino, board [6][10]string) {
	if len(pentaminoesLeft) == 0 {
		// Board is filled send message on channel
		return
	}
	for i, p := range pentaminoesLeft {
		for _, pv := range p.Permutations {
			// place pv
			for y, row := range board {
				for x, _ := range row {
					pos := Vector2{x, y}
					if canPlacePentamino(board, pv, pos) {
						placePentamino2(pv, board, p.Id, pos)

						pl := append(pentaminoesLeft[:i], pentaminoesLeft[i+1:]...)
						start(pl, board)
						removePentamino(pv, board, pos)
					}
				}
			}
		}
	}
	// pentamino := pentaminoesLeft[0]
	// startPentamino(pentamino.Permutations, board)
	// start(pentaminoesLeft[1:], board)
}
func boardToString(board [6][10]string) string {
	str := ""
	for _, row := range board {
		for _, col := range row {
			str += col
		}
	}
	return str
}
func removePentamino(p [5]Vector2, board [6][10]string, pos Vector2) {
	for _, v := range p {
		board[v.X][v.Y] = "."
	}
}
func placePentamino2(p [5]Vector2, board [6][10]string, id string, pos Vector2) {
	for _, v := range p {
		board[v.X][v.Y] = id
	}
}

// func placePentamino2(permutationsLeft [][5]Vector2, pId string) {
// 	for i, pv := range permutationsLeft {
// 		//place pv,
// 		// pvl := append(permutationsLeft[:i], permutationsLeft[i+1:]...)
// 		// placePentamino2(pvl, pId)
// 	}
// }

func validateBoard(board [6][10]string) bool {

	for y, row := range board {
		for x, col := range row {
			noEmptySpaces := getEmptySpaces(x, y, board)
			if noEmptySpaces%5 == 0 {
				continue
			}
			return false
		}
	}

	return true
}

func getEmptySpaces(x int, y int, board [6][10]string) int {
	if board[y][x] != "." {
		return 0
	}

}

func startPentamino(permutationsLeft [][5]Vector2, board [6][10]string) {

	if len(permutationsLeft) == 1 {
		return
	}

	startPentamino(permutationsLeft[1:], board)
}

func placePentamino(board [6][10]string, p Pentamino, pos Vector2) {
	for _, pp := range p.Permutations {
		canPlace := canPlacePentamino(board, pp, pos)
		if canPlace {
			for _, pv := range pp {
				x := pos.X + pv.X
				y := pos.Y + pv.Y
				board[y][x] = p.Id
			}
		}
	}
}

func canPlacePentamino(board [6][10]string, p [5]Vector2, pos Vector2) bool {
	for _, pv := range p {
		x := pos.X + pv.X
		y := pos.Y + pv.Y
		if x < 0 || y < 0 || x > 9 || y > x {
			return false
		}
		if board[y][x] == "." {
			continue
		} else {
			return false
		}
	}
	return true
}

func printBoard(board [6][10]string) {
	var boardToPrint string
	for _, row := range board {
		for _, col := range row {
			boardToPrint += col // + "\t"
		}
		boardToPrint += "\n"
	}
	fmt.Println(boardToPrint)
}

func generatePentaminoes() []Pentamino {
	xPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {2, 0}, {1, 1}, {1, -1}},
	}
	xPentamino := Pentamino{Id: "V", Permutations: xPermutations}

	iPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
	}
	iPentamino := Pentamino{Id: "I", Permutations: iPermutations}

	zPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {1, 0}, {2, -1}, {2, 0}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {2, -2}},
	}
	zPentamino := Pentamino{Id: "Z", Permutations: zPermutations}

	vPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		// Following rotations are not needed
		// {{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
		// {{0, 0}, {1, 0}, {2, -2}, {2, -1}, {2, 0}},
		// {{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
	}
	vPentamino := Pentamino{Id: "V", Permutations: vPermutations}

	tPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {2, 1}},
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {2, 0}},
		{{0, 0}, {1, 0}, {2, -1}, {2, 0}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 0}},
	}

	tPentamino := Pentamino{Id: "T", Permutations: tPermutations}

	wPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 0}, {1, -1}, {1, 0}, {2, -2}, {2, -1}},
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 0}, {0, 1}, {1, -1}, {1, 0}, {2, -1}},
	}

	wPentamino := Pentamino{Id: "W", Permutations: wPermutations}

	uPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}},
		{{0, 0}, {0, 2}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {2, 1}},
	}

	uPentamino := Pentamino{Id: "U", Permutations: uPermutations}

	lPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}},
		{{0, 0}, {1, 0}, {2, 0}, {3, -1}, {3, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}},
		{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {3, 1}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 0}},
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {3, 1}},
		{{0, 0}, {1, -3}, {1, -2}, {1, -1}, {1, 0}},
	}

	lPentamino := Pentamino{Id: "L", Permutations: lPermutations}

	nPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {1, -2}, {1, -1}, {1, 0}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {3, 1}},
		{{0, 0}, {0, 1}, {0, 2}, {1, -1}, {1, 0}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {3, 1}},
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {1, 3}},
		{{0, 0}, {1, 0}, {2, -1}, {2, 0}, {3, -1}},
		{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 3}},
		{{0, 0}, {1, -1}, {1, 0}, {2, -1}, {3, -1}},
	}

	nPentamino := Pentamino{Id: "N", Permutations: nPermutations}

	yPermutations := [][5]Vector2{
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {1, 1}},
		{{0, 0}, {1, -1}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 1}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {3, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 2}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 0}, {3, 0}},
		{{0, 0}, {1, -1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 0}, {1, 0}, {2, -1}, {2, 0}, {3, 0}}}

	yPentamino := Pentamino{Id: "Y", Permutations: yPermutations}

	fPermutations := [][5]Vector2{
		{{0, 0}, {1, -1}, {1, 0}, {1, 1}, {2, 1}},
		{{0, 0}, {0, 1}, {1, -1}, {1, 0}, {2, 0}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, -1}, {2, 0}},
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {2, -1}},
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 0}, {1, -1}, {1, 0}, {1, 1}, {2, -1}},
		{{0, 0}, {1, -1}, {1, 0}, {2, 0}, {2, 1}}}

	fPentamino := Pentamino{Id: "F", Permutations: fPermutations}

	pPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 1}},
		{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 0}, {2, 1}},
		{{0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 0}, {1, -1}, {1, 0}, {2, -1}, {2, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}}}

	pPentamino := Pentamino{Id: "P", Permutations: pPermutations}

	return []Pentamino{xPentamino, iPentamino, zPentamino, vPentamino, tPentamino, wPentamino, uPentamino, lPentamino, nPentamino, yPentamino, fPentamino, pPentamino}
}
