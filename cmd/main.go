package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/CaptainFallaway/BestSchedulingAlgo/internal"
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
	tm := graphing.NewTerminalManager(true)

	inputComp := internal.InputBox{}
	fpsComp := internal.FpsBox{}
	textSaving := internal.SavedText{}
	diagram1 := internal.Diagram{Out: &textSaving}
	// testingComp := internal.Testing{}

	tm.Row().Col(&inputComp, 5).Col(&fpsComp)
	tm.Row(4).Col(&diagram1).Col(&diagram1).Col(&textSaving)

	// The render loop
	go func() {
		c := 0
		t := time.Now()

		for {
			tm.Render()
			c++

			if time.Since(t) > time.Second {
				fpsComp.Fps = c
				c = 0
				t = time.Now()
			}
		}
	}()

	// The main loop / input loop
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

			textSaving.AddRow(input)
			diagram1.SetValues(input)

			inputComp.Clear()
		}
	}
}
