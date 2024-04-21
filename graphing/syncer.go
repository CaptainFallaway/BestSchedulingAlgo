package graphing

import "sync"

type ISyncer interface {
	Done()
}

type Syncer struct {
	wg sync.WaitGroup
}

func (s *Syncer) Start(delta int, c *chan TermPixel) {
	s.wg.Add(delta)
	go s.loop(c)
}

func (s *Syncer) Done() {
	// Inline sync done call, removing overhead of a function call
	s.wg.Add(-1)
}

func (s *Syncer) loop(c *chan TermPixel) {
	s.wg.Wait()
	close(*c)
}
