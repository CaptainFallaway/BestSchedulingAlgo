package main

import (
	"os"

	"github.com/CaptainFallaway/BestSchedulingAlgo/terminal"
	"github.com/CaptainFallaway/BestSchedulingAlgo/utils"
	"github.com/nathan-fiscaletti/consolesize-go"
)

var (
	w = 0
	h = 0
)

var Border = utils.Borders{
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomLeft:  "└",
	BottomRight: "┘",
	Horizontal:  "─",
	Vertical:    "│",
}

func main() {
	os.Stdout.Sync()

	terminal.ClearScreen()
	for {
		width, height := consolesize.GetConsoleSize()
		if w != width || h != height {
			terminal.ClearScreen()
			renderBox(width, height)
			w = width
			h = height
			terminal.CursorToPos(2, height-2)
		}
	}
}

func renderBox(cols, rows int) {
	terminal.HideCursor()

	vectors := make([]utils.TermVector, (cols*rows)-((cols-1)+(rows-1)))

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			char, ok := getChar(c, r, cols, rows)
			if ok {
				vectors = append(vectors, utils.TermVector{X: c, Y: r, Char: char})
			}
		}
	}

	utils.RenderVectors(vectors)

	terminal.ShowCursor()
}

func getChar(c, r, cols, rows int) (string, bool) {
	// Check if the pos is in the top left corner
	if c == 1 && r == 1 {
		return Border.TopLeft, true
	}
	// Check if the pos is in the top right corner
	if c == cols-1 && r == 1 {
		return Border.TopRight, true
	}
	// Check if the pos is in the bottom left corner
	if c == 1 && r == rows-1 {
		return Border.BottomLeft, true
	}
	// Check if the pos is in the bottom right corner
	if c == cols-1 && r == rows-1 {
		return Border.BottomRight, true
	}
	// Check if the pos is on the border
	if r == 1 || r == rows-1 {
		return Border.Horizontal, true
	}
	// Check if the pos is on the border
	if c == 1 || c == cols-1 {
		return Border.Vertical, true
	}
	return "", false
}
