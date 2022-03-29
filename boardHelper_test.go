package main

import (
	"testing"
)

func TestValidateBoard(t *testing.T) {
	board := [6][10]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"x", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{"x", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	res := ValidateBoard(board)
	if res == true {
		t.Fatalf("Board with 2 impossible squares was still valid")
	}
}

func TestValidateBoard2(t *testing.T) {
	board := [6][10]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"x", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "x", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	res := ValidateBoard(board)
	if res == false {
		t.Fatalf("Board with valid spaces was false")
	}
}

func TestValidateBoardFirstHafOkSecondHalfNot(t *testing.T) {
	board := [6][10]string{
		{".", ".", ".", ".", "x", "x", "x", "x", "x", "x"},
		{".", ".", ".", ".", "x", ".", ".", ".", ".", "."},
		{".", ".", ".", "x", "x", "x", ".", ".", ".", "."},
		{".", ".", ".", "x", "x", ".", ".", ".", ".", "."},
		{".", ".", ".", "x", "x", ".", ".", ".", ".", "."},
		{".", ".", ".", "x", "x", ".", ".", ".", ".", "."},
	}
	res := ValidateBoard(board)
	if res == true {
		t.Fatalf("Board with invalid spaces was true")
	}
}
