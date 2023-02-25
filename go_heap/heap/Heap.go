package heap

import (
	"fmt"
)

type Position int

type Heap struct {
	area   []int
	tail   Position
	bottom Position
}

type HeapError struct {
	code int
	msg  string
}

func NewHeap(size Position) *Heap {
	return &Heap{
		area:   make([]int, size),
		tail:   0,
		bottom: size,
	}
}

func (e HeapError) Error() string {
	return fmt.Sprintln("ERROR [%d]: %s", e.code, e.msg)
}

func NewHeapError(code int, msg string) HeapError {
	return HeapError{
		code: code,
		msg:  msg,
	}
}

func (h *Heap) Parent(child Position) (Position, error) {
	if child <= 0 || child >= h.tail {
		return 0, NewHeapError(100, "No parent")
	}
	return (child - 1) / 2, nil
}

func (h *Heap) leftChild(parent Position) (Position, error) {
	left := 2*parent + 1
	if left >= h.tail {
		return 0, NewHeapError(200, "No left child")
	}
	return left, nil
}

func (h *Heap) rightChild(parent Position) (Position, error) {
	right := 2*parent + 2
	if right >= h.tail {
		return 0, NewHeapError(200, "No right child")
	}
	return right, nil
}

func (h *Heap) PercolateDown(parent Position) {
	var min Position
	l, lerr := h.leftChild(parent)
	if lerr == nil && h.area[l] < h.area[parent] {
		min = l
	} else {
		min = parent
	}
	r, rerr := h.rightChild(parent)
	if rerr == nil && h.area[r] < h.area[parent] {
		min = r
	} else {
		min = parent
	}
	if min != parent {
		h.area[parent], h.area[min] = h.area[min], h.area[parent]
	}
	h.PercolateDown(min)
}

func (h *Heap) Add(value int) error {
	if h.tail == h.bottom {
		return NewHeapError(800, "Heap area full")
	}
	pos := h.tail
	h.tail++
	for {
		if pos == 0 {
			break
		}
		cur := (pos - 1) / 2
		if value <= h.area[cur] {
			break
		}
		h.area[pos] = h.area[cur]
		pos = cur
	}
	h.area[pos] = value
	return nil
}

func (h *Heap) Remove(pos Position) (int, error) {
	if h.tail == 0 {
		return 0, NewHeapError(900, "No heap")
	}
	v := h.area[0]
	h.area[0] = h.area[h.tail-1]
	h.tail--
	h.PercolateDown(0)
	return v, nil
}

func (h *Heap) Get() (int, error) {
	if h.tail == 0 {
		return 0, NewHeapError(900, "No heap")
	}
	return h.area[0], nil
}

func (h *Heap) Show() {
	for k, v := range h.area {
		if k >= int(h.tail) {
			break
		}
		fmt.Printf("%v,", v)
	}
	fmt.Println("")
}
