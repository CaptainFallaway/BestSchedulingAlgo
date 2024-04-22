package graphing

import (
	"os"
	"strings"
	"time"

	"github.com/nathan-fiscaletti/consolesize-go"
)

type TerminalManger struct {
	Layout

	// Previous size of the terminal
	width  int
	height int

	termBuffer Buffer
	renderSync Syncer

	delta         int64
	previousFrame time.Time
}

func NewTerminalManager() *TerminalManger {
	hideCursor()
	clearScreen()

	width, height := consolesize.GetConsoleSize()

	return &TerminalManger{
		Layout:     Layout{},
		width:      width,
		height:     height,
		termBuffer: *NewBuffer(width, height),
		renderSync: Syncer{},
	}
}

func (tm *TerminalManger) Render() {
	width, height := consolesize.GetConsoleSize()
	pixelChannel := make(chan termPixel, width*height)

	if width != tm.width || height != tm.height {
		tm.Layout.CalcSizes(width, height)
		tm.width = width
		tm.height = height
		tm.termBuffer = *NewBuffer(width, height)
		clearScreen()
	}

	if len(tm.Layout.Items) != tm.Layout.ItemsCount {
		tm.Layout.CalcSizes(width, height)
		tm.termBuffer = *NewBuffer(width, height)
		clearScreen()
	}

	tm.renderSync.start(len(tm.Layout.Items), &pixelChannel)

	for _, item := range tm.Layout.Items {
		go item.Renderable.Render(tm.delta, item.Dimensions, constructChanSendFunc(pixelChannel, item.Bounds), &tm.renderSync)
	}

	instructions := strings.Builder{}

	for tp := range pixelChannel {
		buffItem := tm.termBuffer.Get(tp.X, tp.Y)

		if tp != buffItem && !(buffItem.Char == 0 && tp.Char == ' ') {
			tm.termBuffer.Set(tp)
			instructions.WriteString(tp.ToAnsi())
		}
	}

	os.Stdout.WriteString(instructions.String())

	tm.delta = time.Since(tm.previousFrame).Milliseconds()
	tm.previousFrame = time.Now()
}
