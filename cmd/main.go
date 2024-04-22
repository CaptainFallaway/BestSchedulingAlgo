package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
	"github.com/CaptainFallaway/BestSchedulingAlgo/internal"
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

	tm := graphing.NewTerminalManager()

	tm.Row().Col(&internal.DiagramBox{})
	tm.Row(2).Col(&internal.DiagramBox{}, 2).Col(&internal.DiagramBox{})
	tm.Row(3).Col(&internal.DiagramBox{}, 1).Col(&internal.DiagramBox{}, 4)

	go func() {
		for {
			tm.Render()
		}
	}()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if char == 'q' {
			break
		}
	}
}
