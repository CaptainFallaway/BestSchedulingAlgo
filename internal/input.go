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

const proompt = "   Enter Process: "

func (t *InputBox) Render(delta int64, ts graphing.CompDimensions, ps graphing.PixelSender, s graphing.ISyncer) {
	defer s.Done()

	// Render the border
	for r := 0; r < ts.Height; r++ {
		for c := 0; c < ts.Width; c++ {
			border := getBorder(c, r, ts.Width, ts.Height, ' ')
			if border != ' ' {
				ps(border, c, r, graphing.FgWhite, graphing.Bold)
			}
		}
	}

	inputted := t.getInput()

	proomptRunes := []rune(proompt)
	proomptLen := len(proomptRunes)

	halfSize := ts.Height / 2
	if ts.Height%2 == 0 {
		halfSize--
	}

	for col := 1; col < ts.Width-1; col++ {
		if col-1 < proomptLen {
			ps(proomptRunes[col-1], col, halfSize, graphing.FgWhite, graphing.Bold)
			continue
		}

		if col-proomptLen-1 < len(inputted) && len(inputted) != 0 {
			if col-proomptLen-1 == t.cursorPos {
				ps(inputted[col-proomptLen-1], col, halfSize, graphing.FgWhite, graphing.Bold, graphing.BgBlue)
			} else {
				ps(inputted[col-proomptLen-1], col, halfSize, graphing.FgWhite, graphing.Bold)
			}
		} else {
			if col-proomptLen-1 == t.cursorPos {
				ps('█', col, halfSize, graphing.FgBlue, graphing.Bold)
			} else {
				ps(' ', col, halfSize)
			}
		}
	}
}
