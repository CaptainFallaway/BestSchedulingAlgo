package graphing

// TSize is a supposed to be a struct parsed to a components render function
type TSize struct {
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

// This is not final solution, but it is a start
// Just returning the computed width for each component
func getSizes(elements, maxWidth, maxHeight int) []TSize {
	sizes := make([]TSize, elements)

	width := maxWidth / elements
	height := maxHeight
	for i := 0; i < elements; i++ {
		sizes[i] = TSize{width + 1, height + 1, width * i, 0} // offset Y is always 0 for now
	}

	return sizes
}
