package main

import (
	"fmt"
	"github.com/cabouffard/mot_cache/game"
)

type HiddenWords struct {
	board *game.Board
}

func NewHiddenWords(size int) *HiddenWords {
	board := game.NewBoard(size)

	return &HiddenWords{board: board}
}

func (hw *HiddenWords) PlayGame() {
	fmt.Printf("GAME HAS STARTED! \n")
	for hw.board.NbFreeSpace() > 40 {
		orientation := hw.board.SelectOrientation()
		wordLength := hw.board.SelectWordLength()
		x, y := hw.board.SelectWordPosition(orientation, wordLength)
		// println()
		// fmt.Printf("x: %d, y: %d\n", x+1, y+1)
		// fmt.Printf("orientation: %v \n", orientation)
		// fmt.Printf("length: %v \n", wordLength)
		query := hw.board.FindQuery(x, y, orientation, wordLength)
		foundWord, err := hw.board.FindWord(query)
		if err != nil {
			// fmt.Printf("%v \n", err)
			continue
		}
		word := game.NewWord(foundWord, x, y, orientation)

		// fmt.Printf("word: %v \n", word)
		hw.board.SetWord(word)
	}
	fmt.Printf("%v \n", hw.board)
	hw.board.PrintListWords()
}

func play() {
	game := NewHiddenWords(16)
	game.PlayGame()
}

func main() {
	if true {
		play()
	}
}
