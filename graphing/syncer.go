package graphing

import "sync"

type ISyncer interface {
	Done()
}

type Syncer struct {
	wg sync.WaitGroup
	c  *chan TermPixel
}

func NewSyncer(delta int, c *chan TermPixel) ISyncer {
	t := Syncer{
		wg: sync.WaitGroup{},
		c:  c,
	}

	t.wg.Add(delta)
	go t.loop()
	return &t
}

func (s *Syncer) Done() {
	s.wg.Done()
}

func (s *Syncer) loop() {
	s.wg.Wait()
	close(*s.c)
}
