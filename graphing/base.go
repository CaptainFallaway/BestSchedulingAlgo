package graphing

import "sync"

type Renderable interface {
	Render(*RenderedComponent, *sync.WaitGroup)
}

type RenderInstruction struct {
	Char string
	X    int
	Y    int
}

type ComponentAccumulator struct {
	Content []RenderInstruction
	Width   int
	Height  int
}
