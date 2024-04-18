package graphing

type Renderable interface {
	Render(TSize, *chan TermPixel, ISyncer)
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
