package graphing

import "math/rand"

type DiagramBox struct{}

var chars = []rune{' ', '░', '▒', '▓', '█'}

func getRandomChar() rune {
	return chars[rand.Intn(15)%5]
}

func (d *DiagramBox) Render(size TSize, ts chan TermPixel, syncer ISyncer) {
	defer syncer.Done()

	var char rune
	var randChar rune

	for r := 1; r < size.Height; r++ {
		for c := 1; c < size.Width; c++ {
			randChar = getRandomChar()
			char = getBorder(c, r, size.Width, size.Height, randChar)
			ts <- TermPixel{Char: char, X: c + size.OffsetX, Y: r + size.OffsetY}
		}
	}
}
