package main

import (
	"fmt"
)

func main() {
	// board := NewTestBoard()
	board := NewBoard(16)

	fmt.Printf("STARTS \n")

	// for i := 0; i < 500; i++ {
	for board.NbFreeSpace() > 30 {
		orientation := board.SelectOrientation()
		wordLength := board.SelectWordLength()
		x, y := board.SelectWordPosition(orientation, wordLength)
		// println()
		// fmt.Printf("x: %d, y: %d\n", x+1, y+1)
		// fmt.Printf("orientation: %v \n", orientation)
		// fmt.Printf("length: %v \n", wordLength)
		query := board.FindQuery(x, y, orientation, wordLength)
		foundWord, err := board.FindWord(query)
		if err != nil {
			// fmt.Printf("%v \n", err)
			continue
		}
		word := NewWord(foundWord, x, y, orientation)

		// fmt.Printf("word: %v \n", word)
		board.SetWord(word)
	}
	fmt.Printf("%v \n", board)
	board.printListWords()
}
