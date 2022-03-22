package main

type Pentamino struct {
	Id string
	// Different rotations/reflection of the same Pentomino
	Permutations [][5]Vector2
}
