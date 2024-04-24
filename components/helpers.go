package components

import "github.com/CaptainFallaway/BestSchedulingAlgo/graphing"

type Stack[T graphing.AnsiOption | int] struct {
	Arr []T

	cycle bool
	idx   int
}

func NewColorStack() *Stack[graphing.AnsiOption] {
	return &Stack[graphing.AnsiOption]{
		Arr: []graphing.AnsiOption{
			graphing.FgBrightBlue,
			graphing.FgBrightRed,
			graphing.FgBrightYellow,
			graphing.FgBrightGreen,
			graphing.FgBrightCyan,
			graphing.FgBrightMagenta,
		},
		cycle: true,
	}
}

func (c *Stack[T]) Pop() T {
	if len(c.Arr) == 0 {
		var ret T
		return ret
	}

	if c.idx >= len(c.Arr) {
		c.idx = 0
	}

	// Remember c.idx is initialized as 0
	ret := c.Arr[c.idx]

	if !c.cycle {
		c.Arr = c.Arr[1:]
	} else {
		c.idx++
	}

	return ret
}

type Borders struct {
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune
	Horizontal  rune
	Vertical    rune
}

// ═║╔╗╚╝
// ┌┐└┘─│
var Border = Borders{
	TopLeft:     '┌',
	TopRight:    '┐',
	BottomLeft:  '└',
	BottomRight: '┘',
	Horizontal:  '─',
	Vertical:    '│',
}

func getBorder(c, r, width, height int, char rune) rune {
	if c == 0 && r == 0 {
		return Border.TopLeft
	}
	if c == width-1 && r == 0 {
		return Border.TopRight
	}
	if c == 0 && r == height-1 {
		return Border.BottomLeft
	}
	if c == width-1 && r == height-1 {
		return Border.BottomRight
	}
	if r == 0 || r == height-1 {
		return Border.Horizontal
	}
	if c == 0 || c == width-1 {
		return Border.Vertical
	}
	return char
}
