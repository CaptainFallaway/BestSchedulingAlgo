package main

import (
	"os"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/CaptainFallaway/BestSchedulingAlgo/utils"
	"github.com/eiannone/keyboard"
)

func main() {
	os.Stdout.Sync()
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	utils.HideCursor()
	utils.ClearScreen()

	tm := graphing.NewTerminalManager()

	tm.AddComponent(&graphing.DiagramBox{})
	tm.AddComponent(&graphing.DiagramBox{})

	for {
		tm.Render()
		tm.Flush()
	}
}
