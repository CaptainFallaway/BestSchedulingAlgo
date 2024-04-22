package graphing

import "sync"

type ISyncer interface {
	Done()
}

type Syncer struct {
	wg sync.WaitGroup
}

func (s *Syncer) start(delta int, c *chan termPixel) {
	s.wg.Add(delta)
	go s.waiter(c)
}

// When the render function is done, call this to decrement to sync all the components goroutines
func (s *Syncer) Done() {
	// Inline sync done call, removing overhead of a function call
	s.wg.Add(-1)
}

func (s *Syncer) waiter(c *chan termPixel) {
	s.wg.Wait()
	close(*c)
}
