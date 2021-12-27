package prioq

import (
	"constraints"
	"errors"
)

type CompareFunc[T any] func(a T, b T) bool

// Heap is a representation of a binary heap data structure
type Heap[T any] struct {
	size     int
	compare  func(a T, b T) bool
	elements []T
}

// New ...
func New[T constraints.Ordered](elements []T, initialCapacity int) *Heap[T] {
	return NewWithCompareFunc(elements, initialCapacity, func(a T, b T) bool {
		return a < b
	})
}

// NewWithCompareFunc ...
func NewWithCompareFunc[T any](elements []T, initialCapacity int, cf CompareFunc[T]) *Heap[T] {
	elems := make([]T, len(elements), initialCapacity)
	copy(elems, elements)

	h := &Heap[T]{
		size:     len(elems),
		compare:  cf,
		elements: elems,
	}

	h.heapify()

	return h
}

// heapify makes a heap of the slice in-place
func (h *Heap[T]) heapify() {
	i := h.size/2 - 1
	for i >= 0 {
		left := 2*i + 1
		right := left + 1

		if right > h.size-1 {
			// Look at only the left child
			if h.compare(h.elements[i], h.elements[left]) {
				h.elements[i], h.elements[left] = h.elements[left], h.elements[i]
			}
		} else {
			// Look at both the left and right child
			rightIsLarger := h.compare(h.elements[left], h.elements[right])
			var compareIndex int

			if rightIsLarger {
				compareIndex = right
			} else {
				compareIndex = left
			}

			shouldSwap := h.compare(h.elements[i], h.elements[compareIndex])
			if shouldSwap {
				h.elements[i], h.elements[compareIndex] = h.elements[compareIndex], h.elements[i]
				if compareIndex < h.size/2 {
					i = compareIndex + 1
				}
			}
		}
		i--
	}
}

// Insert adds a new element to the heap
// The time complexity is  O(log(n)), n = # of elements in the heap
func (h *Heap[T]) Insert(x T) {
	h.elements = append(h.elements, x)
	h.size++

	// Fix the heap
	i := h.size - 1
	for i >= 0 && h.compare(h.elements[parent(i)], h.elements[i]) {
		h.elements[parent(i)], h.elements[i] = h.elements[i], h.elements[parent(i)]
		i = parent(i)
	}
}

// Extract returns the element at the root of the heap
// The time complexity is  O(log(n)), n = # of elements in the heap
func (h *Heap[T]) Extract() (T, error) {
	if h.size == 0 {
		// Trick for getting a generic zero value
		var t T
		return t, errors.New("heap: empty, no element to extract")
	}

	h.elements[h.size-1], h.elements[0] = h.elements[0], h.elements[h.size-1]
	removedElem := h.elements[h.size-1]

	h.size--

	// Only one node left, no need to fix the heap
	if h.size == 1 {
		return removedElem, nil
	}

	// Fix the heap
	i := 0
	for i < h.size-1 {
		child := h.largerChild(i)

		if h.compare(h.elements[i], h.elements[child]) {
			h.elements[i], h.elements[child] = h.elements[child], h.elements[i]
			if child < h.size/2-1 {
				i = child
			}
		} else {
			break
		}
	}

	return removedElem, nil
}

// IsEmpty indicates if the heap has no elements left
func (h *Heap[T]) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap[T]) largerChild(i int) int {
	left := 2*i + 1
	right := left + 1

	if right > h.size-1 {
		return left
	}

	if h.compare(h.elements[left], h.elements[right]) {
		return right
	}
	return left
}

func parent(i int) int {
	return (i - 1) / 2
}
