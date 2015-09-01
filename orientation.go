package main

type Orientation int

const (
	S Orientation = 1 + iota
	E
	N
	W
	NE
	NW
	SE
	SW
)

var orientations = [...]string{
	"South ↓",
	"East →",
	"North ↑",
	"West ←",
	// "Northeast ↗",
	"Nortwest ↖",
	"Southeast ↘",
	// "Southwest ↙",
}

func (o Orientation) String() string { return orientations[o-1] }
