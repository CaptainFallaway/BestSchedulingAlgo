package internal

import "github.com/CaptainFallaway/BestSchedulingAlgo/graphing"

type Stack[T graphing.AnsiOption | int] struct {
	Arr []T
}

func NewColorStack() *Stack[graphing.AnsiOption] {
	return &Stack[graphing.AnsiOption]{
		Arr: []graphing.AnsiOption{
			graphing.FgBlue,
			graphing.FgRed,
			graphing.FgGreen,
			graphing.FgYellow,
			graphing.FgMagenta,
			graphing.FgCyan,
			graphing.FgWhite,
		},
	}
}

func (c *Stack[T]) Pop() T {
	if len(c.Arr) == 0 {
		var ret T
		return ret
	}

	ret := c.Arr[0]
	c.Arr = c.Arr[1:]
	return ret
}
