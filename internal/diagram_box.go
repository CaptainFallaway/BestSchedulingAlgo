package internal

import (
	"math/rand"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

type DiagramBox struct {
	something int
	target    int
	climax    bool
	ticks     int
}

func (d *DiagramBox) Render(size graphing.CompDimensions, ps graphing.PixelSender, syncer graphing.ISyncer) {
	defer syncer.Done()

	var char rune
	if d.something == d.target && d.climax {
		d.target = rand.Intn(size.Height - 1)
	}

	for r := 0; r < size.Height; r++ {
		for c := 0; c < size.Width; c++ {
			if r <= d.something {
				char = '█'
			} else {
				char = ' '
			}
			ps(getBorder(c, r, size.Width, size.Height, char), c, r)
		}
	}

	if d.something < d.target {
		if d.ticks == 200 {
			d.something++
			d.ticks = 0
		}
	} else {
		if d.ticks == 200 {
			d.something--
			d.climax = true
			d.ticks = 0
		}
	}

	d.ticks++
}