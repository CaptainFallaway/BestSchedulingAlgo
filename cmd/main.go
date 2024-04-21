package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/CaptainFallaway/BestSchedulingAlgo/internal"
	"github.com/CaptainFallaway/BestSchedulingAlgo/utils"
	"github.com/eiannone/keyboard"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	os.Stdout.SyscallConn()

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	utils.HideCursor()
	utils.ClearScreen()

	tm := graphing.NewTerminalManager()

	tm.AddRow()

	tm.AddRenderable(&internal.DiagramBox{}, 0, 3)
	tm.AddRenderable(&internal.DiagramBox{}, 0)
	tm.AddRenderable(&internal.DiagramBox{}, 0)

	tm.AddRow(3)

	tm.AddRenderable(&internal.DiagramBox{}, 1)

	tm.AddRow()

	tm.AddRenderable(&internal.DiagramBox{}, 2)

	tm.AddRow()

	tm.AddRenderable(&internal.DiagramBox{}, 3)

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

		if key == keyboard.KeyEsc {
			tm.Layout.AddRenderable(&internal.DiagramBox{}, 3, 2)
		}

		if char == 'q' {
			break
		}
	}
}
