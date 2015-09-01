package main

import (
	"fmt"
)

const GridNilValue rune = '_'

type Board struct {
	grid  [][]rune
	size  int
	words []string
	orm   *ORM
}

func NewBoard(size int) *Board {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
	}

	for x, row := range grid {
		for y, _ := range row {
			grid[x][y] = GridNilValue
		}
	}

	orm := InitDatabase("./words.db")
	orm = orm.Reset()
	return &Board{grid: grid, size: size, orm: orm}
}

func NewTestBoard() *Board {
	board := NewBoard(5)
	board.orm.Reset()

	board.grid[0][0] = 'a'
	board.grid[0][1] = 'a'
	board.grid[0][2] = 'a'

	return board
}

func (board *Board) SelectWordPosition(orientation Orientation, wordLength int) (int, int) {
	x := random(0, board.size)
	y := random(0, board.size)
	if orientation == S || orientation == N || orientation == SE || orientation == NW {
		totalLength := y + wordLength
		for totalLength > board.size {
			y--
			totalLength = y + wordLength
		}
	}

	if orientation == E || orientation == W || orientation == SE || orientation == NW {
		totalLength := x + wordLength
		for totalLength > board.size {
			x--
			totalLength = x + wordLength
		}
	}

	return x, y
}

func (board *Board) SelectOrientation() Orientation {
	nbOrientations := len(orientations)
	return Orientation(random(1, nbOrientations+1))
}

func (board *Board) SelectWordLength() int {
	length := random(3, board.size)
	return length
}

func (board *Board) Set(x, y int, value rune) {
	board.grid[y][x] = value
}

func (board *Board) FindQuery(x, y int, orientation Orientation, wordLength int) string {
	var query string = ""
	if orientation == S || orientation == N {
		for i := y; i < y+wordLength; i++ {
			char := board.Get(x, i)
			if *char == GridNilValue {
				query = query + "_"
			} else {
				s := fmt.Sprintf("%c", *char)
				query = query + s
			}
			if orientation == N {
				query = reverse(query)
			}
		}
	}

	// if orientation == NE || orientation == SW {
	// 	for i := 0; i < length; i++ {
	// 		y := y + i
	// 		x := x + i
	//
	// 		char := board.Get(x, y)
	// 		if *char == GridNilValue {
	// 			query = query + "_"
	// 		} else {
	// 			s := fmt.Sprintf("%c", *char)
	// 			query = query + s
	// 		}
	// 		if orientation == N {
	// 			query = reverse(query)
	// 		}
	//
	// 		if orientation == SW {
	// 			query = reverse(query)
	// 	}
	// }

	if orientation == SE || orientation == NW {
		for i := 0; i < wordLength; i++ {
			y := y + i
			x := x + i

			char := board.Get(x, y)
			if *char == GridNilValue {
				query = query + "_"
			} else {
				s := fmt.Sprintf("%c", *char)
				query = query + s
			}
			if orientation == N {
				query = reverse(query)
			}
		}
	}

	if orientation == E || orientation == W {
		for i := x; i < x+wordLength; i++ {
			char := board.Get(i, y)
			if *char == GridNilValue {
				query = query + "_"
			} else {
				s := fmt.Sprintf("%c", *char)
				query = query + s
			}
			if orientation == W {
				query = reverse(query)
			}
		}
	}
	return query

}

func (board *Board) FindWord(query string) (string, error) {
	foundWord, err := board.orm.FindWord(query)
	if err != nil {
		return "", err
	}

	testCount := 0
	for stringInSlice(foundWord, board.words) {
		testCount += 1
		foundWord, err = board.orm.FindWord(query)
		if err != nil {
			return "", err
		}
		if testCount == 3 {
			err := fmt.Errorf("Unable to word that has not been used with the query: %v", query)
			return "", err
		}

	}
	return foundWord, nil
}
func (board *Board) SetWord(word *Word) {
	if word.orientation == S || word.orientation == N {
		for i := word.y; i < word.y+word.length; i++ {
			value := word.value
			if word.orientation == N {
				value = reverse(word.value)
			}
			r := rune([]rune(value)[i-word.y])
			board.Set(word.x, i, r)
		}
	}

	// if word.orientation == NE || word.orientation == SW {
	// 	for i := 0; i < word.length; i++ {
	// 		y := word.y + i
	// 		x := word.x + i
	// 		value := word.value
	// 		if word.orientation == SW {
	// 			value = reverse(word.value)
	// 		}
	// 		r := rune([]rune(value)[i])
	// 		board.Set(x, y, r)
	// 	}
	// }

	if word.orientation == SE || word.orientation == NW {
		for i := 0; i < word.length; i++ {
			y := word.y + i
			x := word.x + i
			value := word.value
			if word.orientation == NW {
				value = reverse(word.value)
			}
			r := rune([]rune(value)[i])
			board.Set(x, y, r)
		}
	}

	if word.orientation == E || word.orientation == W {
		for i := word.x; i < word.x+word.length; i++ {
			value := word.value
			if word.orientation == W {
				value = reverse(word.value)
			}
			r := rune([]rune(value)[i-word.x])
			board.Set(i, word.y, r)
		}
	}
	board.words = append(board.words, word.value)
}

func (board *Board) printListWords() {
	println("List of words: ")
	for _, word := range board.words {
		println(word)
	}
}

func (board *Board) Get(x, y int) *rune {
	return &board.grid[y][x]
}

func (board *Board) NbFreeSpace() int {
	count := 0
	for _, row := range board.grid {
		for _, value := range row {
			if value == GridNilValue {
				count += 1
			}
		}
	}
	return count

}

func (board *Board) String() string {
	// For each rows
	var str string
	for _, row := range board.grid {
		// For each columns
		for _, value := range row {
			str = str + fmt.Sprintf("[%+q]", value)
		}
		str = str + "\n"
	}
	return str
}
