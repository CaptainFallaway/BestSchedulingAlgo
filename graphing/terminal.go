package graphing

import (
	"os"
	"strings"
	"time"

	terminal "github.com/CaptainFallaway/BestSchedulingAlgo/utils"
	"github.com/nathan-fiscaletti/consolesize-go"
)

type TerminalManger struct {
	Layout

	// Previous size of the terminal
	Width  int
	Height int

	TermBuffer Buffer
	RenderSync Syncer

	delta         int64
	previousFrame time.Time
}

func NewTerminalManager() *TerminalManger {
	width, height := consolesize.GetConsoleSize()
	return &TerminalManger{
		Layout:     Layout{},
		Width:      width,
		Height:     height,
		TermBuffer: *NewBuffer(width, height),
		RenderSync: Syncer{},
	}
}

func (tm *TerminalManger) Render() {
	width, height := consolesize.GetConsoleSize()
	pixelChannel := make(chan TermPixel, width*height)

	if width != tm.Width || height != tm.Height {
		tm.Layout.CalcSizes(width, height)
		tm.Width = width
		tm.Height = height
		tm.TermBuffer = *NewBuffer(width, height)
		terminal.ClearScreen()
	}

	if len(tm.Layout.Items) != tm.Layout.ItemsCount {
		tm.Layout.CalcSizes(width, height)
		tm.TermBuffer = *NewBuffer(width, height)
		terminal.ClearScreen()
	}

	tm.RenderSync.Start(len(tm.Layout.Items), &pixelChannel)

	for _, item := range tm.Layout.Items {
		go item.Renderable.Render(tm.delta, item.Dimensions, constructChanSendFunc(pixelChannel, item.Bounds), &tm.RenderSync)
	}

	instructions := strings.Builder{}

	for tp := range pixelChannel {
		buffItem := tm.TermBuffer.Get(tp.X, tp.Y)

		if tp != buffItem && !(buffItem.Char == 0 && tp.Char == ' ') {
			tm.TermBuffer.Set(tp)
			instructions.WriteString(tp.ToAnsi())
		}
	}

	os.Stdout.WriteString(instructions.String())

	tm.delta = time.Since(tm.previousFrame).Milliseconds()
	tm.previousFrame = time.Now()
}
