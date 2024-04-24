package terminal

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

	termBuffer buffer
	renderSync syncer // the syncer for the parallel / concurrent component rendering

	// The delta stuff
	delta         int64
	previousFrame time.Time
	fps           uint16

	stopChan chan bool

	sleepPerRenderCycle time.Duration // Just a int64, uhm, weird implementation
}

func NewTerminalManager(options ...TerminalOption) *TerminalManger {
	width, height := consolesize.GetConsoleSize()

	temp := &TerminalManger{
		Layout:     Layout{},
		width:      width,
		height:     height,
		termBuffer: *NewBuffer(width, height),
		renderSync: syncer{},
	}

	for _, option := range options {
		option(temp)
	}

	return temp
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
			instructions.WriteString(tp.toAnsi())
		}
	}

	os.Stdout.WriteString(instructions.String())

	tm.delta = time.Since(tm.previousFrame).Milliseconds()
	tm.previousFrame = time.Now()

	if tm.sleepPerRenderCycle > 0 {
		time.Sleep(tm.sleepPerRenderCycle * time.Millisecond)
	}
}

func (tm *TerminalManger) GetFps() uint16 {
	return tm.fps
}

func (tm *TerminalManger) Start() {
	tm.stopChan = make(chan bool)

	hideCursor()
	clearScreen()

	go func() {
		var c uint16
		t := time.Now()

		for {
			tm.Render()
			c++

			select {
			case <-tm.stopChan:
				return
			default:
				if time.Since(t).Seconds() >= 1 {
					tm.fps = c
					t = time.Now()
					c = 0
				}

				continue
			}

		}
	}()
}

func (tm *TerminalManger) Stop() {
	tm.stopChan <- true
}
