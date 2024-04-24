package graphing

import (
	"fmt"
	"os"
	"strings"
)

func constructChanSendFunc(pc chan termPixel, size componentBounds) func(rune, int, int, ...AnsiOption) {
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

		pc <- termPixel{Char: char, X: x + size.OffsetX + 1, Y: y + size.OffsetY + 1, ansiOpts: sb.String()}
	}
}

func getSpan(span ...int) int {
	if len(span) >= 1 {
		return span[0]
	}

	return 1
}

func clearScreen() {
	os.Stdout.WriteString("\x1b[2J")
}

func hideCursor() {
	os.Stdout.WriteString("\x1b[?25l")
}
