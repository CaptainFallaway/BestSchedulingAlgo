package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/CaptainFallaway/BestSchedulingAlgo/internal"
	"github.com/eiannone/keyboard"
)

func main() {
	// This is for profiling the application
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	tm := graphing.NewTerminalManager()

	inputComp := internal.InputBox{}

	tm.Row().Col(&inputComp)
	tm.Row(4).Col(&internal.DiagramBox{}).Col(&internal.DiagramBox{}).Col(&internal.DiagramBox{})

	go func() {
		for {
			tm.Render()
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
		}
	}
}
