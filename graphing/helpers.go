package graphing

import (
	"fmt"
	"time"
)

func constructChanSendFunc(pc chan TermPixel, size componentBounds) func(rune, int, int) {
	return func(char rune, x, y int) {
		if x < 0 || y < 0 || x >= size.Width || y >= size.Height {
			panic(
				fmt.Sprintf("Pixel out of bounds for one of the components: %d, %d, %d, %d",
					x,
					y,
					size.Width,
					size.Height,
				),
			) // TODO: Think of some way presenting the object that caused the error
		}

		pc <- TermPixel{Char: char, X: x + size.OffsetX + 1, Y: y + size.OffsetY + 1}
	}
}

func timeNow() float64 {
	return float64(time.Now().UnixNano())
}

// Helper function since i do the exact same thing in two places
func getSpan(span ...int) int {
	if len(span) >= 1 {
		return span[0]
	}

	return 1
}
