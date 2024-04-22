package internal

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

type Diagram struct {
	Values []float64
	m      sync.RWMutex

	Out  *SavedText
	Prev string
}

func (d *Diagram) SetValues(inpt string) {
	sNums := strings.Split(inpt, " ")
	nums := make([]float64, 0, len(sNums))

	for _, sNum := range sNums {
		num, err := strconv.ParseFloat(sNum, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	d.m.Lock()
	defer d.m.Unlock()

	d.Values = nums
}

func (d *Diagram) GetValues() []float64 {
	d.m.RLock()
	defer d.m.RUnlock()

	return d.Values
}

func sum(values []float64) float64 {
	var sum float64

	for _, v := range values {
		sum += v
	}

	return sum
}

// Förändrings faktor
func getFfs(values []float64, sumValues float64, dunno int) []int {
	arr := make([]int, 0, len(values))

	for _, val := range values {
		x := (val / sumValues) * float64(dunno)
		arr = append(arr, int(x))
	}

	return arr
}

func (d *Diagram) Render(delta int64, size graphing.CompDimensions, ps graphing.PixelSender, syncer graphing.ISyncer) {
	defer syncer.Done()

	// // Render the border
	// for r := 0; r < size.Height; r++ {
	// 	for c := 0; c < size.Width; c++ {
	// 		char := getBorder(c, r, size.Width, size.Height, ' ')
	// 		if char != ' ' {
	// 			ps(char, c, r)
	// 		}
	// 	}
	// }

	// Will need to clean this up
	colorStack := NewColorStack()

	values := d.GetValues()
	sumValues := sum(values)
	rowCounts := getFfs(values, sumValues, (size.Height-2)*(size.Width-1))
	countStack := Stack[int]{
		Arr: rowCounts,
	}

	// Some weird debug stuff
	x := fmt.Sprint(rowCounts) + " " + fmt.Sprint(sumValues) + " " + fmt.Sprint((size.Height-2)*(size.Width-1))
	if x != d.Prev {
		d.Out.AddRow(x)
	}
	d.Prev = x

	counted := 1
	color := colorStack.Pop()
	count := countStack.Pop()

	for r := 1; r < size.Height-1; r++ {
		for c := 1; c < size.Width-1; c++ {
			if counted >= count {
				counted = 0
				count = countStack.Pop()
				color = colorStack.Pop()
			}

			ps('█', c, r, color)

			counted++
		}
	}
}
