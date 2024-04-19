package main

import (
	"os"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/CaptainFallaway/BestSchedulingAlgo/utils"
	"github.com/eiannone/keyboard"
)

func main() {
	os.Stdout.SyscallConn()

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	utils.HideCursor()
	utils.ClearScreen()

	tm := graphing.NewTerminalManager()

	tm.AddComponent(&graphing.DiagramBox{})

	go func() {
		for {
			tm.Render()
		}
	}()

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			panic(err)
		}

		if char == 'a' {
			tm.AddComponent(&graphing.DiagramBox{})
		}

		if char == 'd' {
			tm.Components = tm.Components[:len(tm.Components)-1]
		}
	}
}
