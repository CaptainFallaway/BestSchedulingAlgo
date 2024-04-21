package internal

import (
	"fmt"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

type Testing struct {
	Fps int
}

func (t *Testing) Render(ts graphing.CompDimensions, ps graphing.PixelSender, s graphing.ISyncer) {
	defer s.Done()

	x := fmt.Sprint(t.Fps)

	for c, char := range x {
		ps(char, c, 0)
	}
}
