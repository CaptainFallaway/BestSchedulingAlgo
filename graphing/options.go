package graphing

import "time"

type TerminalOption func(*TerminalManger)

// Interesting pattern AnthonyGG has taught me

func FrameSleep(sleepMs int64) TerminalOption {
	return func(tm *TerminalManger) {
		tm.sleepPerRenderCycle = time.Duration(sleepMs)
	}
}
