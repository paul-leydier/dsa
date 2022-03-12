package src

import "testing"

func TestMaxHeap_Insert(t *testing.T) {
	// Insert a number larger than all existing in a MaxHeap, and check the heapify process
	heap := MaxHeap[int]{
		elements: []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7},
		length:   11,
	}
	heap.Insert(63)
	expectedSlice := []int{63, 16, 50, 14, 8, 48, 20, 9, 1, 5, 7, 34}
	for i := 0; i < len(heap.elements); i++ {
		if expectedSlice[i] != heap.elements[i] {
			t.Fatalf("element %d is unexpected: got %d, should have gotten %d", i, heap.elements[i], expectedSlice[i])
		}
	}
	if heap.length != 12 {
		t.Fatalf("length attribute is unexpected: got %d, expected %d", heap.length, 3)
	}
}

func TestMaxHeap_Extract(t *testing.T) {
	// Insert a number larger than all existing in a MaxHeap, and check the heapify process
	heap := MaxHeap[int]{
		elements: []int{63, 16, 50, 14, 8, 48, 20, 9, 1, 5, 7, 34},
		length:   12,
	}
	result := heap.Extract()
	expectedSlice := []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7}
	if result != 63 {
		t.Fatalf("result is unexpected: got %d, expected %d", result, 63)
	}
	for i := 0; i < len(heap.elements); i++ {
		if expectedSlice[i] != heap.elements[i] {
			t.Fatalf("element %d is unexpected: got %d, should have gotten %d", i, heap.elements[i], expectedSlice[i])
		}
	}
	if heap.length != 11 {
		t.Fatalf("length attribute is unexpected: got %d, expected %d", heap.length, 3)
	}
}
