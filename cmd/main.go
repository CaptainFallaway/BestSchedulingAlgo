package main

import (
	"fmt"

	"github.com/CaptainFallaway/BestSchedulingAlgo/utils"
)

var (
	w = 0
	h = 0
)

var Border = utils.Borders{
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomLeft:  "└",
	BottomRight: "┘",
	Horizontal:  "─",
	Vertical:    "│",
}

type Node struct {
	Next *Node
	Val  int
}

func main() {
	// os.Stdout.Sync()

	// terminal.ClearScreen()
	// for {
	// 	width, height := consolesize.GetConsoleSize()

	// 	if w != width || h != height {
	// 		terminal.ClearScreen()
	// 		w = width
	// 		h = height
	// 		renderBox(width, height)
	// 		terminal.CursorToPos(2, height-2)
	// 	}
	// }

	var root *Node = &Node{Val: 0}
	copy := root
	for i := 1; i < 11; i++ {
		new := &Node{Val: i}
		root.Next = new
		root = new
	}

	recursiveIteration(copy)
}

func recursiveIteration(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	recursiveIteration(node.Next)
}
