package terminal

// \x1b[2J - Clear screen: Clears the entire screen.
// \x1b[0;0H - Move cursor to the top-left corner.
// \x1b[K - Clear line: Clears the current line from the cursor position to the end of the line.
// \x1b[2K - Clear line: Clears the entire line.
// \x1b[A - Move cursor up one line.
// \x1b[B - Move cursor down one line.
// \x1b[C - Move cursor forward (right) one position.
// \x1b[D - Move cursor backward (left) one position.
// \x1b[<n>A - Move cursor up n lines.
// \x1b[<n>B - Move cursor down n lines.
// \x1b[<n>C - Move cursor forward n positions.
// \x1b[<n>D - Move cursor backward n positions.
// \x1b[<x>;<y>H - Move cursor to position <x>,<y>.
// \x1b[<n>m - Set text attributes like color, style, etc. (e.g., \x1b[31m for red text).

import "fmt"

func write(s string, a ...any) {
	fmt.Printf(s, a...)
}

func ClearScreen() {
	write("\x1b[2J")
}

func ClearToEnd() {
	write("\x1b[K")
}

func ClearWholeLine() {
	write("\x1b[2K")
}

func ClearLine(y int) {
	write("\033[%vH\033[J", y)
}

func CursorToPos(x, y int) {
	write("\x1b[%v;%vH", y, x)
}

func CursorUp(ammount ...int) {
	if len(ammount) > 0 {
		write("\x1b[%vA", ammount[0])
	} else {
		write("\x1b[A")
	}
}

func CursorDown(ammount ...int) {
	if len(ammount) > 0 {
		write("\x1b[%vB", ammount[0])
	} else {
		write("\x1b[B")
	}
}

func CursorRight(ammount ...int) {
	if len(ammount) > 0 {
		write("\x1b[%vC", ammount[0])
	} else {
		write("\x1b[C")
	}
}

func CursorLeft(ammount ...int) {
	if len(ammount) > 0 {
		write("\x1b[%vD", ammount[0])
	} else {
		write("\x1b[D")
	}
}

func SaveCursorPos() {
	write("\033[S")
}

func RestoreCursorPos() {
	write("\033[U")
}

func HideCursor() {
	write("\033[?25l")
}

func ShowCursor() {
	write("\033[?25h")
}
