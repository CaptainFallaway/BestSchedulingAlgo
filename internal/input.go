package internal

import (
	"sync"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

type InputBox struct {
	cursorPos int
	inputted  []rune
	m         sync.RWMutex
}

func (t *InputBox) Insert(r rune) {
	t.m.Lock()
	defer t.m.Unlock()

	at := t.cursorPos
	t.inputted = append(t.inputted[:at], append([]rune{r}, t.inputted[at:]...)...)

	t.cursorPos++
}

func (t *InputBox) Backspace() {
	t.m.Lock()
	defer t.m.Unlock()

	if len(t.inputted) == 0 || t.cursorPos == 0 {
		return
	}

	at := t.cursorPos - 1
	t.inputted = append(t.inputted[:at], t.inputted[at+1:]...)

	t.cursorPos--
}

func (t *InputBox) CursorLeft() {
	t.m.Lock()
	defer t.m.Unlock()
	if t.cursorPos > 0 {
		t.cursorPos--
	}
}

func (t *InputBox) CursorRight() {
	t.m.Lock()
	defer t.m.Unlock()
	if t.cursorPos < len(t.inputted) {
		t.cursorPos++
	}
}

func (t *InputBox) Home() {
	t.m.Lock()
	defer t.m.Unlock()
	t.cursorPos = 0
}

func (t *InputBox) End() {
	t.m.Lock()
	defer t.m.Unlock()
	t.cursorPos = len(t.inputted)
}

func (t *InputBox) getInput() []rune {
	t.m.RLock()
	defer t.m.RUnlock()
	return t.inputted
}

func (t *InputBox) Render(delta int64, ts graphing.CompDimensions, ps graphing.PixelSender, s graphing.ISyncer) {
	defer s.Done()

	inputted := t.getInput()

	halfSize := ts.Height / 2

	for r := 0; r < ts.Height; r++ {
		for c := 0; c < ts.Width; c++ {

			if r != halfSize {
				ps(' ', c, r, graphing.FgWhite, graphing.Bold)
				continue
			}

			if len(inputted) > c && len(inputted) > 0 {
				if c == t.cursorPos {
					ps(inputted[c], c, r, graphing.FgBlack, graphing.Bold, graphing.BgWhite)
				} else {
					ps(inputted[c], c, r, graphing.FgWhite, graphing.Bold)
				}
			} else {
				if c == t.cursorPos {
					ps('█', c, r, graphing.FgWhite, graphing.Bold)
				} else {
					ps(' ', c, r, graphing.FgWhite, graphing.Bold)
				}
			}
		}
	}
}
