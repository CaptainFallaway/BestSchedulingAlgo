package graphing

import "math/rand"

type DiagramBox struct{}

var chars = []rune{' ', '.'}

func getRandomChar() rune {
	return chars[rand.Intn(3)%2]
}

func (d *DiagramBox) Render(size TSize, rs *chan TermPixel, syncer ISyncer) {
	var char rune
	var randChar rune
	for r := 1; r < size.Height; r++ {
		for c := 1; c < size.Width; c++ {
			randChar = getRandomChar()
			char = getBorder(c, r, size.Width, size.Height, randChar)
			(*rs) <- TermPixel{Char: char, X: c + size.OffsetX, Y: r + size.OffsetY}
		}
	}
	syncer.Done()
}
