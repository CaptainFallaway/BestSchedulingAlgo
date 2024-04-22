package graphing

import (
	"fmt"
	"strings"
	"time"
)

func constructChanSendFunc(pc chan TermPixel, size componentBounds) func(rune, int, int, ...AnsiOption) {
	return func(char rune, x, y int, ansiopts ...AnsiOption) {
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

		sb := strings.Builder{}
		for _, opt := range ansiopts {
			sb.WriteString(string(opt))
		}

		pc <- TermPixel{Char: char, X: x + size.OffsetX + 1, Y: y + size.OffsetY + 1, ansiOpts: sb.String()}
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
