package components

import (
	"math"
	"slices"
	"sync"

	"github.com/CaptainFallaway/BestSchedulingAlgo/terminal"
)

type Diagram struct {
	DiagramName string

	Values []float64
	m      sync.RWMutex

	Labels [][]rune
}

func (d *Diagram) Update(arr []float64, labels []string) {
	d.m.Lock()
	defer d.m.Unlock()

	slices.Sort(arr)

	d.Values = arr

	d.Labels = make([][]rune, len(labels))

	for i, label := range labels {
		d.Labels[i] = []rune(label)
	}
}

func (d *Diagram) GetValues() []float64 {
	d.m.RLock()
	defer d.m.RUnlock()

	return d.Values
}

func (d *Diagram) Render(delta int64, size terminal.CompDimensions, ps terminal.PixelSender, syncer terminal.ISyncer) {
	defer syncer.Done()

	runeName := []rune(d.DiagramName)

	// Render the border
	for r := 0; r < size.Height; r++ {
		for c := 0; c < size.Width; c++ {
			if r == 0 && size.Width-len(d.DiagramName) > 4 && len(d.DiagramName) > 0 && c-2 < len(d.DiagramName) && c-1 > 0 {
				ps(runeName[c-2], c, r, terminal.Bold)
				continue
			}

			char := getBorder(c, r, size.Width, size.Height, ' ')

			if char != ' ' {
				ps(char, c, r, terminal.Bold)
			}
		}
	}

	values := d.GetValues()

	// ac := 1
	// for _, item := range d.Labels {
	// 	for _, char := range item {
	// 		ps(char, ac, 1)
	// 		ac++
	// 	}
	// 	ac++
	// }

	// return

	// May use sync.Pool here
	colorStack := NewColorStack()
	countStack := Stack[int]{
		Arr: getDiagramPixelCounts(values, (size.Height-2)*(size.Width-2)),
	}
	labelStack := Stack[[]rune]{
		Arr: d.Labels,
	}

	// Rendering of the diagram graph
	prev := 0
	counted := 0
	color := colorStack.Pop()
	count := countStack.Pop()
	label := labelStack.Pop()
	co := (count / 2) - count + len(label) - 1

	renderBlank := len(countStack.Arr) == 0

	for r := 1; r < size.Height-1; r++ {
		for c := 1; c < size.Width-1; c++ {
			if counted == count+prev {
				prev = counted
				count = countStack.Pop()
				color = colorStack.Pop()
				label = labelStack.Pop()
				co = (count / 2) - count + len(label) - 1
			}

			if renderBlank {
				ps('█', c, r, terminal.FgBlack, terminal.Bold)
				counted++
				continue
			}

			if count-len(label) > 0 && co >= 0 && co < len(label) && len(label) > 0 {
				ps(label[co], c, r, color.Bg, terminal.Bold)
			} else {
				ps('█', c, r, color.Fg, terminal.Bold)
			}

			co++
			counted++
		}
	}
}

func sum(values []float64) float64 {
	var sum float64

	for _, v := range values {
		sum += v
	}

	return sum
}

// getDiagramPixelCounts returns a slice that contains counts,
// and these counts are the pixels on the screen that represent the value.
func getDiagramPixelCounts(values []float64, frameSize int) []int {
	arr := make([]int, 0, len(values))

	sumValues := sum(values)

	for _, val := range values {
		// get change factor and multiply by the ammount TerminalPixels (Cells)
		x := (val / sumValues) * float64(frameSize)
		arr = append(arr, int(math.Ceil(x)))
	}

	return arr
}
