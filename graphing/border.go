package graphing

import (
	"github.com/CaptainFallaway/BestSchedulingAlgo/utils"
)

var Border = utils.Borders{
	TopLeft:     '┌',
	TopRight:    '┐',
	BottomLeft:  '└',
	BottomRight: '┘',
	Horizontal:  '─',
	Vertical:    '│',
}

func getBorder(c, r, width, height int, char rune) string {
	if c == 1 && r == 1 {
		return Border.TopLeft
	}
	if c == width-1 && r == 1 {
		return Border.TopRight
	}
	if c == 1 && r == width-1 {
		return Border.BottomLeft
	}
	if c == width-1 && r == width-1 {
		return Border.BottomRight
	}
	if r == 1 || r == height-1 {
		return Border.Horizontal
	}
	if c == 1 || c == width-1 {
		return Border.Vertical
	}
	return char
}
