package internal

import (
	"strings"
	"sync"

	"github.com/CaptainFallaway/BestSchedulingAlgo/graphing"
)

const maxRows = 20

type SavedText struct {
	Rows []string
	m    sync.RWMutex
}

func (s *SavedText) AddRow(row string) {
	if strings.TrimSpace(row) == "" {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	if len(s.Rows) > maxRows {
		s.Rows = s.Rows[1:]
	}

	s.Rows = append(s.Rows, row)
}

func (s *SavedText) Clear() {
	s.m.Lock()
	defer s.m.Unlock()
	s.Rows = []string{}
}

func (s *SavedText) getRows() []string {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.Rows
}

func (s *SavedText) Render(delta int64, size graphing.CompDimensions, ps graphing.PixelSender, syncer graphing.ISyncer) {
	defer syncer.Done()

	rows := s.getRows()

	// Render the border
	for r := 0; r < size.Height; r++ {
		for c := 0; c < size.Width; c++ {
			char := getBorder(c, r, size.Width, size.Height, ' ')
			if char != ' ' {
				ps(char, c, r)
			}
		}
	}

	var row []rune

	for r := 1; r < size.Height-1; r++ {
		if r <= len(rows) && len(rows) != 0 {
			row = []rune(rows[r-1])
		} else {
			return
		}
		for c := 1; c < size.Width-1; c++ {
			if c <= len(row) && len(row) != 0 {
				ps(row[c-1], c, r)
			} else {
				ps(' ', c, r)
			}
		}
	}
}
