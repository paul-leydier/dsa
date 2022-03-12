package src

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	elements []T
	length   int
}

// Insert inserts a value within the MaxHeap structure,
// then Heapifies it from the bottom
func (h *MaxHeap[T]) Insert(value T) {
	h.elements = append([]T(h.elements), value)
	h.length++
	h.heapifyFromBottom()
}

// Extract returns the largest value from the MaxHeap structure (the one at the top),
// removing it from the structure, and then Heapifies it from the top
func (h *MaxHeap[T]) Extract() T {
	maxVal := h.elements[0]
	h.elements = h.elements[1:]
	h.length--
	h.heapifyFromTop()
	return maxVal
}

func parent(idx int) int {
	if idx < 0 {
		panic("cannot find parent of idx < 0")
	}
	return (idx - 1) / 2
}

func (h *MaxHeap[T]) heapifyFromBottom() {
	idx := h.length - 1
	p := parent(idx)
	for idx > 0 {
		if h.elements[idx] <= h.elements[p] {
			return
		}
		h.elements[idx], h.elements[p] = h.elements[p], h.elements[idx]
		idx = p
		p = parent(idx)
	}
}

func (h *MaxHeap[T]) heapifyFromTop() {

}
