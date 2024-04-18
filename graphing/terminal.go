package graphing

import (
	"fmt"
	"sync"

	terminal "github.com/CaptainFallaway/BestSchedulingAlgo/utils"
)

type TerminalManger struct {
	Components []IComponent
}

func (tm *TerminalManger) AddComponent(c IComponent) {
	tm.Components = append(tm.Components, c)
}

func (tm *TerminalManger) Render() {
	wg := sync.WaitGroup{}
	rcs := make([]*RenderedComponent, 0, len(tm.Components))

	for _, comp := range tm.Components {
		wg.Add(1)
		temp := new(RenderedComponent)
		rcs = append(rcs, temp)
		go comp.Render(temp, &wg)
	}

	wg.Wait()

	for _, rendered := range rcs {
		for _, vector := range rendered.Content {
			terminal.CursorToPos(vector.X, vector.Y)
			fmt.Printf("%c", vector.Char)
		}
	}
}
