package internal

import (
	"fmt"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

type FpsBox struct {
	Fps int
}

func (f *FpsBox) Render(delta int64, size graphing.CompDimensions, ps graphing.PixelSender, syncer graphing.ISyncer) {
	defer syncer.Done()

	if delta == 0 {
		return
	}

	halfSize := size.Height / 2

	for r := 0; r < size.Height; r++ {
		for c := 0; c < size.Width; c++ {
			char := getBorder(c, r, size.Width, size.Height, ' ')
			if char != ' ' {
				ps(char, c, r)
			}
		}
	}

	fpsRunes := []rune("      FPS: ")

	for i, r := range fpsRunes {
		ps(r, i+1, halfSize)
	}

	for i, r := range []rune(fmt.Sprint(f.Fps)) {
		ps(r, i+len(fpsRunes), halfSize)
	}
}
