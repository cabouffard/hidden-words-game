package game

import (
	"fmt"
)

type Word struct {
	value       string
	length      int
	orientation Orientation
	x, y        int
}

func NewWord(value string, x, y int, orientation Orientation) *Word {
	length := len(value)
	return &Word{value: value, x: x, y: y, orientation: orientation, length: length}
}

func (word *Word) String() string {
	return fmt.Sprintf("%v", word.value)
}
