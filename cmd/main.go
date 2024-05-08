package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/CaptainFallaway/BestSchedulingAlgo/components"
	"github.com/CaptainFallaway/BestSchedulingAlgo/internal"
	"github.com/CaptainFallaway/BestSchedulingAlgo/terminal"
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

func clearScreen() {
	os.Stdout.WriteString("\x1b[2J")
}

func hideCursor() {
	os.Stdout.WriteString("\x1b[?25l")
}

func main() {
	startProfiler()
	startKeyboard()
	clearScreen()
	hideCursor()

	defer keyboard.Close()

	// The terminal stuff
	tm := terminal.NewTerminalManager(terminal.FrameSleep(10))

	fpsBox := components.FpsBox{}
	inputComp := components.InputBox{}
	cpuDiagram1 := components.Diagram{DiagramName: "CPU 1 Fifo"}
	cpuDiagram2 := components.Diagram{DiagramName: "CPU 2 TimeShare"}
	cpuDiagram3 := components.Diagram{DiagramName: "CPU 3 Prio"}

	// Hello Dan
	cpu1 := internal.NewCpu(&internal.Fifo{}, &cpuDiagram1)
	cpu2 := internal.NewCpu(&internal.TimeShare{}, &cpuDiagram2)
	cpu3 := internal.NewCpu(&internal.Prio{}, &cpuDiagram3)
	dispatcher := internal.NewDispatcher(cpu1, cpu2, cpu3)

	tm.Row().Col(&inputComp, 5).Col(&fpsBox)
	tm.Row(4).Col(&cpuDiagram1).Col(&cpuDiagram2).Col(&cpuDiagram3)

	tm.Start()

	history := internal.History{}

	// Fps counter goroutine
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			fpsBox.Fps = int(tm.GetFps())
		}
	}()

	go func() {
		for {
			dispatcher.Work()
			time.Sleep(250 * time.Millisecond)
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
		case keyboard.KeyArrowUp:
			input, ok := history.Up()
			if ok {
				inputComp.SetInput(input)
			}
		case keyboard.KeyArrowDown:
			input, ok := history.Down()
			if ok {
				inputComp.SetInput(input)
			} else {
				inputComp.Clear()
			}
		case keyboard.KeyEnter:
			input := inputComp.GetInput()
			processes := MarshalInput(input)

			dispatcher.Push(processes)

			inputComp.Clear()
			history.Add(input)
		}
	}
}

func MarshalInput(input string) []internal.Process {
	slices := strings.Split(input, "|")

	for i, slice := range slices {
		slices[i] = strings.TrimSpace(slice)
	}

	processes := make([]internal.Process, 0, len(slices))

	for _, slice := range slices {
		parts := strings.Split(slice, " ")

		if len(parts) != 3 {
			continue
		}

		name := parts[0]

		execTime, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			execTime = 1000
		}

		prio, err := strconv.ParseUint(parts[2], 10, 16)
		if err != nil {
			prio = 1
		}

		processes = append(processes, internal.Process{Name: name, ExecTime: execTime, Prio: uint16(prio)})
	}

	return processes
}
