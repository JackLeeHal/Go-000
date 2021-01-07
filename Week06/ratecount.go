package Week06

import (
	"errors"
	"sync"
)

// Window implements an sliding window, use ring buffer inside
type Window struct {
	base       []float64
	start      int
	Len        int
	windowSize int
	m          sync.RWMutex
}

// Create initializes the window with a capacity and an window size
func NewWindow(capacity, windowSize int) (*Window, error) {
	if windowSize > capacity {
		return nil, errors.New("capacity can't be smaller than windowsize")
	}

	return &Window{
		m:          sync.RWMutex{},
		base:       make([]float64, 0, capacity),
		windowSize: windowSize,
	}, nil
}

// Add appends an item to the window
func (w *Window) Add(i float64) {
	w.m.Lock()
	w.add(i)
	w.m.Unlock()
}

// add is like Add, but without locking
func (w *Window) add(i float64) {
	// Move all values to front, if would reach end of base
	if w.start+w.Len+1 > cap(w.base) {
		for j := 0; j < w.Len-1; j++ {
			w.base[j] = w.base[w.start+j+1]
		}
		w.start = 0
		w.Len--
	}

	// Check capacity and append if needed
	if len(w.base) < w.start+w.Len+1 {
		w.base = append(w.base, i)
	} else {
		w.base[w.start+w.Len] = i
	}

	// If window is "full" => Move one
	if w.Len == w.windowSize {
		w.start++
	} else {
		w.Len++
	}
}

// Count returns rate count
func (w *Window) Count() float64 {
	w.m.RLock()
	ret := w.base[w.start : w.start+w.Len]
	w.m.RUnlock()
	var r float64
	for _, v := range ret {
		r += v
	}

	return r
}
