package graphing

import "sync"

type DiagramBox struct {
	BaseBoxComponent
}

func (d *DiagramBox) Render(rc *RenderedComponent, wg *sync.WaitGroup) {
	defer wg.Done()

	temp := make([]Vector, 0, d.Width*d.Height)

	for r := 1; r < d.Height; r++ {
		for c := 1; c < d.Width; c++ {
			char := getBorder(c, r, d.Width, d.Height, 'e')
			temp = append(temp, Vector{char, c, r})
		}
	}

	rc.Content = temp
}
