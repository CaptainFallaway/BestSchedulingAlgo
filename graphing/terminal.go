package graphing

import (
	"fmt"

	"github.com/CaptainFallaway/BestSchedulingAlgo/terminal"
	"github.com/nathan-fiscaletti/consolesize-go"
)

type TerminalManger struct {
	Components  []Renderable
	Width       int
	Height      int
	FrontBuffer Buffer
	BackBuffer  Buffer
}

func NewTerminalManager() *TerminalManger {
	width, height := consolesize.GetConsoleSize()
	return &TerminalManger{
		Components:  []Renderable{},
		Width:       width,
		Height:      height,
		FrontBuffer: *NewBuffer(width, height),
		BackBuffer:  *NewBuffer(width, height),
	}
}

func (tm *TerminalManger) AddComponent(c Renderable) {
	tm.Components = append(tm.Components, c)
}

func CursorToPos(x, y int, char rune) {
	fmt.Printf("\x1b[%d;%dH%c", y, x, char)
}

func (tm *TerminalManger) Render() {
	width, height := consolesize.GetConsoleSize()
	instructionChannel := make(chan TermPixel, width*height)

	tm.BackBuffer = *NewBuffer(width, height)

	sizes := getSizes(len(tm.Components))

	syncer := NewSyncer(len(tm.Components), &instructionChannel)
	for i, c := range tm.Components {
		go c.Render(sizes[i], &instructionChannel, syncer)
	}

	for inst := range instructionChannel {
		tm.BackBuffer.Add(inst)
	}
}

func (tm *TerminalManger) Flush() {
	if len(tm.FrontBuffer.BuffArr) != len(tm.BackBuffer.BuffArr) {
		terminal.ClearScreen()
		for _, tp := range tm.BackBuffer.BuffArr {
			CursorToPos(tp.X, tp.Y, tp.Char)
		}

	} else {
		for i, tp := range tm.BackBuffer.BuffArr {
			if tp != tm.FrontBuffer.BuffArr[i] {
				CursorToPos(tp.X, tp.Y, tp.Char)
			}
		}
	}

	tm.FrontBuffer = tm.BackBuffer
}
