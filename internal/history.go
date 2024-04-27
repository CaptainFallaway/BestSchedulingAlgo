package internal

import (
	"strings"
)

type node struct {
	Data string
	Next *node
	Prev *node
}

type History struct {
	Head     *node
	CurrNode *node
}

func (h *History) Add(input string) {
	input = strings.TrimSpace(input)

	if input == "" {
		return
	}

	if h.Head != nil && h.Head.Data == input {
		h.CurrNode = nil
		return
	}

	newNode := &node{
		Data: input,
		Next: h.Head, // This can be nil
		Prev: nil,
	}

	if h.Head != nil {
		h.Head.Prev = newNode
	}

	h.Head = newNode
	h.CurrNode = nil
}

func (h *History) Up() (string, bool) {
	if h.CurrNode == nil {
		if h.Head != nil {
			h.CurrNode = h.Head
			return h.Head.Data, true
		}

		return "", false
	}

	if h.CurrNode.Next == nil {
		return h.CurrNode.Data, false
	}

	h.CurrNode = h.CurrNode.Next
	return h.CurrNode.Data, true
}

func (h *History) Down() (string, bool) {
	if h.CurrNode == nil || h.CurrNode.Prev == nil {
		h.CurrNode = nil
		return "", false
	}

	h.CurrNode = h.CurrNode.Prev
	return h.CurrNode.Data, true
}
