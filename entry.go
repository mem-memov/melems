package melems

import "fmt"

const maxUint = ^uint(0)

type Entry struct {
	elements []uint
	maxIndex uint
}

func newEntry(elements []uint) *Entry {
	return &Entry {
		elements: elements,
		maxIndex: uint(len(elements)) - 1,
	}
}

func (e *Entry) Set(index uint, value uint) error {
	if index > e.maxIndex {
		return fmt.Errorf("index %d too large, max index %d", index, e.maxIndex)
	}

	e.elements[index] = value
	return nil
}

func (e *Entry) Get(index uint) (uint, error) {
	if index > e.maxIndex {
		return 0, fmt.Errorf("index %d too large, max index %d", index, e.maxIndex)
	}

	return e.elements[index], nil
}

func (e *Entry) Increment(index uint) error {
	if index > e.maxIndex {
		return fmt.Errorf("index %d too large, max index %d", index, e.maxIndex)
	}

	if e.elements[index] == maxUint {
		return fmt.Errorf("no increment is possible at index %d", index)
	}

	e.elements[index]++
	return nil
}

func (e *Entry) Decrement(index uint) error {
	if index > e.maxIndex {
		return fmt.Errorf("index %d too large, max index %d", index, e.maxIndex)
	}

	if e.elements[index] == 0 {
		return fmt.Errorf("no decrement is possible at index %d", index)
	}

	e.elements[index]--
	return nil
}