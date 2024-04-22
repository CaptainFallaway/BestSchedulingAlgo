package internal

type Borders struct {
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune
	Horizontal  rune
	Vertical    rune
}

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
