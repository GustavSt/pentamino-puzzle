package main

import "fmt"

func PrintBoard(board [6][10]string) {
	var boardToPrint string
	for _, row := range board {
		for _, col := range row {
			boardToPrint += col + " "
		}
		boardToPrint += "\n"
	}
	fmt.Println(boardToPrint)
}

func RemovePentamino(p [5]Vector2, board *[6][10]string, pos Vector2) {
	for _, v := range p {
		board[v.Y][v.X] = "."
	}
}
func PlacePentamino(p [5]Vector2, board *[6][10]string, id string, pos Vector2) {
	for _, v := range p {
		x := pos.X + v.X
		y := pos.Y + v.Y
		board[y][x] = id
	}
}

func CanPlacePentamino(board *[6][10]string, p [5]Vector2, pos Vector2) bool {
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

func BoardToString(board [6][10]string) string {
	str := ""
	for _, row := range board {
		for _, col := range row {
			str += col
		}
	}
	return str
}

func ValidateBoard(board *[6][10]string) bool {

	// PrintBoard(board)
	checked := make([]Vector2, 0)
	for y, row := range board {
		for x := range row {
			if contains(checked, Vector2{x, y}) {
				continue
			}
			noEmptySpaces := getEmptySpaces(x, y, *board, &checked)
			if len(noEmptySpaces)%5 == 0 {
				continue
			}
			return false
		}
	}

	return true
}

func getEmptySpaces(x int, y int, board [6][10]string, checked *[]Vector2) []Vector2 {
	if contains(*checked, Vector2{x, y}) {
		return make([]Vector2, 0)
	}
	// fmt.Printf("Checking x: %v, y:%v\n", x, y)
	currSpaces := []Vector2{{x, y}}
	*checked = append(*checked, currSpaces...)
	if board[y][x] != "." {
		return make([]Vector2, 0)
	}
	if x+1 < 10 {
		currSpaces = append(currSpaces, getEmptySpaces(x+1, y, board, checked)...)
	}
	if y+1 < 6 {
		currSpaces = append(currSpaces, getEmptySpaces(x, y+1, board, checked)...)
	}
	if x-1 > -1 {
		currSpaces = append(currSpaces, getEmptySpaces(x-1, y, board, checked)...)
	}
	if y-1 > -1 {
		currSpaces = append(currSpaces, getEmptySpaces(x, y-1, board, checked)...)
	}

	return currSpaces
}
func contains(s []Vector2, v Vector2) bool {
	for _, a := range s {
		if a.X == v.X && a.Y == v.Y {
			return true
		}
	}
	return false
}

func GetAnchorBoards() ([][6][10]string, []Pentamino) {
	pentaminoes := generatePentaminoes()
	boards := [][6][10]string{}
	anchorPoints := []Vector2{
		{1, 1},
		{2, 1},
		{3, 1},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
	}
	xPentamino := pentaminoes[0]
	for _, p := range anchorPoints {
		board := getBoard()
		PlacePentamino(xPentamino.Permutations[0], &board, xPentamino.Id, p)
		boards = append(boards, board)
	}
	return boards, pentaminoes[1:]
}

func getBoard() [6][10]string {
	return [6][10]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
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