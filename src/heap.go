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
	maxVal := h.elements[0]                // take highest value
	h.elements[0] = h.elements[h.length-1] // place the latest value in the top spot
	h.elements = h.elements[:h.length-1]   // remove the last spot
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
	currentNode, firstChild := 0, 1
	for c := h.children(currentNode); len(c) > 0; {
		largestChild := firstChild + ArgMax(c)
		if h.elements[largestChild] > h.elements[currentNode] {
			h.elements[largestChild], h.elements[currentNode] = h.elements[currentNode], h.elements[largestChild]
			currentNode = largestChild
		} else {
			smallestChild := firstChild + ArgMin(c)
			if h.elements[smallestChild] > h.elements[currentNode] {
				h.elements[smallestChild], h.elements[currentNode] = h.elements[currentNode], h.elements[smallestChild]
				currentNode = smallestChild
			}
		}
		c, firstChild = h.children(currentNode), currentNode*2+1
	}
}

func (h *MaxHeap[T]) children(idx int) []T {
	potentialLastChild := idx*2 + 2
	if h.length <= potentialLastChild {
		if h.length <= potentialLastChild-1 {
			return []T{}
		}
		return []T{h.elements[potentialLastChild-1]}
	}
	return []T{h.elements[potentialLastChild-1], h.elements[potentialLastChild]}
}
