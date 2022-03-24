package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

func main() {
	fmt.Println("Pentomino puzzle 6x10")
	boards, pentaminoes := GetAnchorBoards()
	c := make(chan string)
	wg := new(sync.WaitGroup)

	for _, p2 := range pentaminoes {
		fmt.Printf("%v ", p2.Id)
	}
	fmt.Println()

	for i, b := range boards {
		// PrintBoard(b)
		wg.Add(1)
		go func(board [6][10]string, boardId int) {
			defer wg.Done()
			start(pentaminoes, &board, c, boardId)
		}(b, i)
		// fmt.Println("-----------")
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	fmt.Println("Started looking for solutions")
	result := make([]string, 0)
	failed := 0
	success := 0
	fmt.Printf("SolutionsFound: %v. Failed attempts: %v.", success, failed)

	writer1 := uilive.New()
	writer2 := writer1.Newline()
	writer3 := writer1.Newline()
	writer4 := writer1.Newline()
	writer5 := writer1.Newline()
	writer6 := writer1.Newline()
	writer1.Start()

	for r := range c {
		if r[0] == 'b' {
			board := StrToBoard(r[1:])
			for y, row := range board {
				// fmt.Printf("%v\n", RowToString(row))
				if y == 0 {
					fmt.Fprintf(writer1, "%v\n", RowToString(row))
				}
				if y == 1 {
					fmt.Fprintf(writer2, "%v\n", RowToString(row))
				}
				if y == 2 {
					fmt.Fprintf(writer3, "%v\n", RowToString(row))
				}
				if y == 3 {
					fmt.Fprintf(writer4, "%v\n", RowToString(row))
				}
				if y == 4 {
					fmt.Fprintf(writer5, "%v\n", RowToString(row))
				}
				if y == 5 {
					fmt.Fprintf(writer6, "%v\n", RowToString(row))
				}
			}
			continue
		}
		if r == "failed" {
			failed++
			fmt.Printf("\rSolutionsFound: %v. Failed attempts: %v.", success, failed)
			continue
		}
		result = append(result, r)
		success++
		fmt.Printf("\rSolutionsFound: %v. Failed attempts: %v.", success, failed)
	}
	// fmt.Fprintln(writer1, "Finished downloading both files :)")
	// writer1.Stop()
	fmt.Println()
	fmt.Printf("No solutions found: %v\n", success)
}

func start(pentaminoesLeft []Pentamino, board *[6][10]string, c chan<- string, boardId int) {
	// if boardId == 0 {
	// 	for _, p2 := range pentaminoesLeft {
	// 		fmt.Printf("%v ", p2.Id)
	// 	}
	// 	fmt.Println()
	// }
	str := BoardToString(*board)
	if !strings.Contains(str, ".") {
		c <- str
		return
	}
	if len(pentaminoesLeft) == 0 {
		// Board is filled send message on channel
		str2 := BoardToString(*board)
		c <- str2
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
							if boardId == 0 {
								c <- fmt.Sprintf("b%v", BoardToString(*board))
								time.Sleep(time.Second / 10)
							}
							pl1 := make([]Pentamino, len(pentaminoesLeft))
							copy(pl1, pentaminoesLeft)
							pl := append(pl1[:i], pl1[i+1:]...)
							// copy(pl, append(pentaminoesLeft[:i], pentaminoesLeft[i+1:]...))
							// if boardId == 0 {
							// 	fmt.Println("Pentaminoes left")
							// 	for _, p2 := range pentaminoesLeft {
							// 		fmt.Printf("%v ", p2.Id)
							// 	}
							// 	fmt.Println()
							// 	fmt.Println("Pentaminos to send on")
							// 	for _, p1 := range pl {
							// 		fmt.Printf("%v ", p1.Id)
							// 	}
							// 	fmt.Println()
							// }
							start(pl, board, c, boardId)
						}
						RemovePentamino(pv, board, pos)
					}
				}
			}
		}
	}
	// c <- "failed"
}
func remove(p []Pentamino, i int) []Pentamino {
	return append(p[:i], p[i+1:]...)
}
