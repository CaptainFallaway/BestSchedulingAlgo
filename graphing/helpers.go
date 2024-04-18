package graphing

import "github.com/nathan-fiscaletti/consolesize-go"

// TSize is a supposed to be a struct parsed to a components render function
type TSize struct {
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

// This is not final solution, but it is a start
// Just returning the computed width for each component
func getSizes(elements int) []TSize {
	maxWidth, maxHeight := consolesize.GetConsoleSize()

	sizes := make([]TSize, elements)

	width := maxWidth / elements
	height := maxHeight
	for i := 0; i < elements; i++ {
		sizes[i] = TSize{width, height, width * i, 0} // offset Y is always 0 for now
	}

	return sizes
}
