package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/CaptainFallaway/BestSchedulingAlgo/components"
	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/eiannone/keyboard"
)

func startProfiler() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
}

func startKeyboard() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
}

func main() {
	startProfiler()
	startKeyboard()
	defer keyboard.Close()

	// The terminal stuff
	tm := graphing.NewTerminalManager(false)

	fpsBox := components.FpsBox{}
	inputComp := components.InputBox{}
	cpuDiagram1 := components.Diagram{DiagramName: "CPU 1"}
	cpuDiagram2 := components.Diagram{DiagramName: "CPU 2"}
	cpuDiagram3 := components.Diagram{DiagramName: "CPU 3"}

	tm.Row().Col(&inputComp, 5).Col(&fpsBox)
	tm.Row(4).Col(&cpuDiagram1).Col(&cpuDiagram2).Col(&cpuDiagram3)

	// tm.Start()
	// on := true

	// In all reality dont use this way of doing the render loop
	go func() {
		c := 0
		t := time.Now()

		for {
			tm.Render()
			c++

			if time.Since(t).Seconds() >= 1 {
				fpsBox.Fps = c
				t = time.Now()
				c = 0
			}
		}
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case 0:
			inputComp.Insert(char)
		case keyboard.KeyBackspace:
			inputComp.Backspace()
		case keyboard.KeySpace:
			inputComp.Insert(' ')
		case keyboard.KeyArrowLeft:
			inputComp.CursorLeft()
		case keyboard.KeyArrowRight:
			inputComp.CursorRight()
		case keyboard.KeyHome:
			inputComp.Home()
		case keyboard.KeyEnd:
			inputComp.End()
		case keyboard.KeyEnter:
			input := inputComp.GetInput()

			cpuDiagram1.TestSetValues(input)
			cpuDiagram2.TestSetValues(input)
			cpuDiagram3.TestSetValues(input)

			inputComp.Clear()
			// case keyboard.KeyEsc:
			// 	if on {
			// 		tm.Stop()
			// 	} else {
			// 		tm.Start()
			// 	}

			// 	on = !on
		}
	}
}
