package components

import (
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

type Diagram struct {
	DiagramName string

	Values []float64
	m      sync.RWMutex

	Labels [][]rune
}

func (d *Diagram) TestSetValues(inpt string) {
	sNums := strings.Split(inpt, " ")
	nums := make([]float64, 0, len(sNums))

	for _, sNum := range sNums {
		num, err := strconv.ParseFloat(sNum, 64)
		if err != nil {
			num = 0
		}
		nums = append(nums, num)
	}

	d.m.Lock()
	defer d.m.Unlock()

	d.Values = nums
}

func (d *Diagram) UpdateValues(vals []float64) {
	d.m.Lock()
	defer d.m.Unlock()

	d.Values = vals
}

func (d *Diagram) GetValues() []float64 {
	d.m.RLock()
	defer d.m.RUnlock()

	return d.Values
}

func (d *Diagram) Render(delta int64, size graphing.CompDimensions, ps graphing.PixelSender, syncer graphing.ISyncer) {
	defer syncer.Done()

	runeName := []rune(d.DiagramName)

	// Render the border
	for r := 0; r < size.Height; r++ {
		for c := 0; c < size.Width; c++ {
			if r == 0 && size.Width-len(d.DiagramName) > 4 && len(d.DiagramName) > 0 && c-2 < len(d.DiagramName) && c-1 > 0 {
				ps(runeName[c-2], c, r, graphing.Bold)
				continue
			}

			char := getBorder(c, r, size.Width, size.Height, ' ')

			if char != ' ' {
				ps(char, c, r, graphing.Bold)
			}
		}
	}

	values := d.GetValues()

	// May use sync.Pool here
	colorStack := NewColorStack()
	countStack := Stack[int]{
		Arr: getDiagramPixelCounts(values, (size.Height-2)*(size.Width-2)),
	}

	prev := 0
	counted := 0
	color := colorStack.Pop()
	count := countStack.Pop()

	for r := 1; r < size.Height-1; r++ {
		for c := 1; c < size.Width-1; c++ {
			if counted == count+prev {
				prev = counted
				count = countStack.Pop()
				color = colorStack.Pop()
			}

			ps('█', c, r, color, graphing.Bold)

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
