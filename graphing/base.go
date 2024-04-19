package graphing

import "fmt"

type Renderable interface {
	Render(TSize, chan TermPixel, ISyncer)
}

type RenderAcc struct {
	Content []TermPixel
}

func (ra *RenderAcc) Append(r TermPixel) {
	ra.Content = append(ra.Content, r)
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
