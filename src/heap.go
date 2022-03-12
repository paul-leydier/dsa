package src

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	elements []T
	length   int
}

// Insert inserts a value within the MaxHeap structure,
// then Heapifies it from the bottom
func (h *MaxHeap[T]) Insert(value T) {

}

// Extract returns the largest value from the MaxHeap structure (the one at the top),
// removing it from the structure, and then Heapifies it from the top
func (h *MaxHeap[T]) Extract() T {
	return nil
}

func (h *MaxHeap[T]) heapifyFromBottom() {

}

func (h *MaxHeap[T]) heapifyFromTop() {

}
