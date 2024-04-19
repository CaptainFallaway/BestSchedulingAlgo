package graphing

import (
	"fmt"

	"github.com/CaptainFallaway/BestSchedulingAlgo/terminal"
	"github.com/nathan-fiscaletti/consolesize-go"
)

type TerminalManger struct {
	Components []Renderable
	CompCount  int
	Width      int
	Height     int
	TermBuffer Buffer
	RenderSync Syncer
}

func NewTerminalManager() *TerminalManger {
	width, height := consolesize.GetConsoleSize()
	return &TerminalManger{
		Components: []Renderable{},
		CompCount:  0,
		Width:      0,
		Height:     0,
		TermBuffer: *NewBuffer(width, height),
		RenderSync: Syncer{},
	}
}

func (tm *TerminalManger) AddComponent(c Renderable) {
	tm.Components = append(tm.Components, c)
}

func (tm *TerminalManger) Render() {
	if len(tm.Components) == 0 {
		return
	}

	width, height := consolesize.GetConsoleSize()
	pixelChannel := make(chan TermPixel, width*height)

	if width != tm.Width || height != tm.Height {
		tm.Width = width
		tm.Height = height
		tm.TermBuffer = *NewBuffer(width, height)
		terminal.ClearScreen()
	}

	if len(tm.Components) != tm.CompCount {
		tm.CompCount = len(tm.Components)
		tm.TermBuffer = *NewBuffer(width, height)
		terminal.ClearScreen()
	}

	sizes := getSizes(len(tm.Components), width, height)

	tm.RenderSync.Start(len(tm.Components), &pixelChannel)

	for i, c := range tm.Components {
		go c.Render(sizes[i], pixelChannel, &tm.RenderSync)
	}

	instructions := ""

	for tp := range pixelChannel {
		indx := ((tp.Y - 1) * width) + (tp.X - 1)
		if tp != tm.TermBuffer.BuffArr[indx] {
			tm.TermBuffer.BuffArr[indx] = tp
			instructions += tp.ToAnsi()
		}
	}

	fmt.Print(instructions) // Might want to use Stdout.Write immediately instead
}
