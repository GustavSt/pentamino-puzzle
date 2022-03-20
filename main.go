package main

import "fmt"

func main() {
	fmt.Println("hello")
	firstVector := Vector2{5, 7}
	fmt.Printf("vecotr: %v\n", firstVector)
}

func generatePentaminoes() []Pentamino {
	xPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {2, 0}, {1, 1}, {1, -1}},
	}
	xPentamino := Pentamino{Permutations: xPermutations}

	iPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
	}
	iPentamino := Pentamino{Permutations: iPermutations}

	zPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {1, 0}, {2, -1}, {2, 0}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {2, -2}},
	}
	zPentamino := Pentamino{Permutations: zPermutations}

	vPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		// Following rotations are not needed
		// {{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
		// {{0, 0}, {1, 0}, {2, -2}, {2, -1}, {2, 0}},
		// {{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
	}
	vPentamino := Pentamino{Permutations: vPermutations}

	tPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {2, 1}},
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {2, 0}},
		{{0, 0}, {1, 0}, {2, -1}, {2, 0}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 0}},
	}

	tPentamino := Pentamino{Permutations: tPermutations}

	wPermutations := [][5]Vector2{
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 0}, {1, -1}, {1, 0}, {2, -2}, {2, -1}},
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 0}, {0, 1}, {1, -1}, {1, 0}, {2, -1}},
	}

	wPentamino := Pentamino{Permutations: wPermutations}

	uPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}},
		{{0, 0}, {0, 2}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {2, 1}},
	}

	uPentamino := Pentamino{Permutations: uPermutations}

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

	lPentamino := Pentamino{Permutations: lPermutations}

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

	nPentamino := Pentamino{Permutations: nPermutations}

	yPermutations := [][5]Vector2{
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {1, 1}},
		{{0, 0}, {1, -1}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 1}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {3, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 2}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 0}, {3, 0}},
		{{0, 0}, {1, -1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 0}, {1, 0}, {2, -1}, {2, 0}, {3, 0}}}

	yPentamino := Pentamino{Permutations: yPermutations}

	fPermutations := [][5]Vector2{
		{{0, 0}, {1, -1}, {1, 0}, {1, 1}, {2, 1}},
		{{0, 0}, {0, 1}, {1, -1}, {1, 0}, {2, 0}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, -1}, {2, 0}},
		{{0, 0}, {1, -2}, {1, -1}, {1, 0}, {2, -1}},
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 0}, {1, -1}, {1, 0}, {1, 1}, {2, -1}},
		{{0, 0}, {1, -1}, {1, 0}, {2, 0}, {2, 1}}}

	fPentamino := Pentamino{Permutations: fPermutations}

	pPermutations := [][5]Vector2{
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 1}},
		{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 0}, {2, 1}},
		{{0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 0}, {1, -1}, {1, 0}, {2, -1}, {2, 0}},
		{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}}}

	pPentamino := Pentamino{Permutations: pPermutations}

	return []Pentamino{xPentamino, iPentamino, zPentamino, vPentamino, tPentamino, wPentamino, uPentamino, lPentamino, nPentamino, yPentamino, fPentamino, pPentamino}
}
