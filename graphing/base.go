package graphing

import (
	"fmt"
)

type PixelSender func(rune, int, int, ...AnsiOption)

type AnsiOption string

type Renderable interface {
	// This method is called every cycle to render the component
	// The arguments should look like:
	// Render(delta int64, size CompDimensions, ps PixelSender, syncer ISyncer)
	Render(int64, CompDimensions, PixelSender, ISyncer)
}

// A components dimensions passed to a objects render function
type CompDimensions struct {
	Width  int
	Height int
}

// A components position and size on the terminal screen
type componentBounds struct {
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

// In the future pixels should either be ANSI instructions with pos
// and with or with other ANSI attributes
type TermPixel struct {
	Char     rune
	X        int
	Y        int
	ansiOpts string
}

func (tp TermPixel) ToAnsi() string {
	return fmt.Sprintf("%s\x1b[%d;%dH%c\x1b[0m", tp.ansiOpts, tp.Y, tp.X, tp.Char)
}

const (
	Bold            AnsiOption = "\x1b[1m"
	Faint           AnsiOption = "\x1b[2m"
	Italic          AnsiOption = "\x1b[3m"
	Underline       AnsiOption = "\x1b[4m"
	SwapFgBg        AnsiOption = "\x1b[7m"
	StrikeThrough   AnsiOption = "\x1b[9m"
	FgBlack         AnsiOption = "\x1b[30m"
	FgRed           AnsiOption = "\x1b[31m"
	FgGreen         AnsiOption = "\x1b[32m"
	FgYellow        AnsiOption = "\x1b[33m"
	FgBlue          AnsiOption = "\x1b[34m"
	FgMagenta       AnsiOption = "\x1b[35m"
	FgCyan          AnsiOption = "\x1b[36m"
	FgLightGray     AnsiOption = "\x1b[37m"
	FgDarkGray      AnsiOption = "\x1b[90m"
	FgBrightRed     AnsiOption = "\x1b[91m"
	FgBrightGreen   AnsiOption = "\x1b[92m"
	FgBrightYellow  AnsiOption = "\x1b[93m"
	FgBrightBlue    AnsiOption = "\x1b[94m"
	FgBrightMagenta AnsiOption = "\x1b[95m"
	FgBrightCyan    AnsiOption = "\x1b[96m"
	FgWhite         AnsiOption = "\x1b[97m"
	BgBlack         AnsiOption = "\x1b[40m"
	BgRed           AnsiOption = "\x1b[41m"
	BgGreen         AnsiOption = "\x1b[42m"
	BgYellow        AnsiOption = "\x1b[43m"
	BgBlue          AnsiOption = "\x1b[44m"
	BgMagenta       AnsiOption = "\x1b[45m"
	BgCyan          AnsiOption = "\x1b[46m"
	BgLightGray     AnsiOption = "\x1b[47m"
	BgDarkGray      AnsiOption = "\x1b[100m"
	BgBrightRed     AnsiOption = "\x1b[101m"
	BgBrightGreen   AnsiOption = "\x1b[102m"
	BgBrightYellow  AnsiOption = "\x1b[103m"
	BgBrightBlue    AnsiOption = "\x1b[104m"
	BgBrightMagenta AnsiOption = "\x1b[105m"
	BgBrightCyan    AnsiOption = "\x1b[106m"
	BgWhite         AnsiOption = "\x1b[107m"
)
