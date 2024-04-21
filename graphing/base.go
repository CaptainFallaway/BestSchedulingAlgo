package graphing

import "fmt"

type PixelSender func(rune, int, int)

type Renderable interface {
	Render(CompDimensions, PixelSender, ISyncer)
}

// A components dimensions passed to a objects render function
type CompDimensions struct {
	Width  int
	Height int
}

// A components position on the screen
type componentBounds struct {
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

// In the future pixels should either be ANSI instructions with pos
// and with or with other ANSI attributes
type TermPixel struct {
	Char rune
	X    int
	Y    int
}

func (tp TermPixel) ToAnsi() string {
	return fmt.Sprintf("\x1b[%d;%dH%c", tp.Y, tp.X, tp.Char)
}
