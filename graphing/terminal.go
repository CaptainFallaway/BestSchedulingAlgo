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

	termBuffer buffer
	renderSync syncer // the syncer for the parallel / concurrent component rendering

	delta         int64
	previousFrame time.Time

	// Thousand FPS mode
	thousandFpsMode bool

	stopChan chan bool
}

func NewTerminalManager(thousandFpsMode bool) *TerminalManger {
	hideCursor()
	clearScreen()

	width, height := consolesize.GetConsoleSize()

	return &TerminalManger{
		Layout:          Layout{},
		width:           width,
		height:          height,
		termBuffer:      *NewBuffer(width, height),
		renderSync:      syncer{},
		thousandFpsMode: thousandFpsMode,
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
			instructions.WriteString(tp.toAnsi())
			// os.Stdout.WriteString(tp.toAnsi())
			// time.Sleep(time.Nanosecond)
		}
	}

	os.Stdout.WriteString(instructions.String())

	tm.delta = time.Since(tm.previousFrame).Milliseconds()
	tm.previousFrame = time.Now()

	if !tm.thousandFpsMode {
		time.Sleep(10 * time.Millisecond)
	}
}

func (tm *TerminalManger) Start() {
	tm.stopChan = make(chan bool)

	go func() {
		for {
			tm.Render()
			select {
			case <-tm.stopChan:
				return
			default:
				continue
			}
		}
	}()
}

func (tm *TerminalManger) Stop() {
	tm.stopChan <- true
}
