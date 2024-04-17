package main

import (
	"fmt"
	"github.com/nathan-fiscaletti/consolesize-go"
)

var border = [6]string{"─", "│", "┌", "┐", "└", "┘"}

func main() {
	cols, rows := consolesize.GetConsoleSize()

	renderBorder(cols, rows)
}

func renderBorder(cols int, rows int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 || r == rows-1 {
				switch c {
					case 0:
						if r == 0 {
							fmt.Print(border[2])
						} else {
							fmt.Print(border[4])
						}
					case cols-1:
						if r == 0 {
							fmt.Print(border[3])
						} else {
							fmt.Print(border[5])
						}
					default:
						fmt.Print(border[0])
				}
			} else {
				if c == 0 || c == cols-1 {
					fmt.Print("")
					fmt.Print(border[1])
				}
			}
			// Nice
		}
		fmt.Print("\n")
	}
}
