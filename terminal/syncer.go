package terminal

import "sync"

// Defer the Done function to sync the
type ISyncer interface {
	Done()
}

type syncer struct {
	wg sync.WaitGroup
}

func (s *syncer) start(delta int, c *chan termPixel) {
	s.wg.Add(delta)
	go s.waiter(c)
}

// When the render function is done, call this to decrement to sync all the components goroutines
func (s *syncer) Done() {
	// Inline sync done call, removing overhead of a function call
	s.wg.Add(-1)
}

func (s *syncer) waiter(c *chan termPixel) {
	s.wg.Wait()
	close(*c)
}
