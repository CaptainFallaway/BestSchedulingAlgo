package utils

import (
	"fmt"

	"github.com/CaptainFallaway/BestSchedulingAlgo/terminal"
)

type Borders struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
}

type TermVector struct {
	X    int
	Y    int
	Char string
}

func RenderVectors(vectors []TermVector) {
	for _, v := range vectors {
		terminal.CursorToPos(v.X, v.Y)
		fmt.Print(v.Char)
	}
}
