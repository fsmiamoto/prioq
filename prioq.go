package prioq

import (
	"constraints"
	"errors"
)

// CompareFunc is a generic function that compares two values and that should return true
// whenever those values should be swapped
type CompareFunc[T any] func(a T, b T) bool

// PrioQ represents a generic priority queue data structure
type PrioQ[T any] struct {
	size     int
	compare  func(a T, b T) bool
	elements []T
}

// New creates a new priority queue for elements in the Ordered constraint.
func New[T constraints.Ordered](elements []T) *PrioQ[T] {
	return NewWithCompareFunc(elements, func(a T, b T) bool {
		return a > b
	})
}

// NewWithCompareFunc creates a new priority queue with the given initial capacity.
// Specifiing an initial capacity can be useful to avoid reallocations but in general
// you can just specify len(elements).
func NewWithCompareFunc[T any](elements []T, cf CompareFunc[T]) *PrioQ[T] {
	elems := make([]T, len(elements))
	copy(elems, elements)

	h := &PrioQ[T]{
		size:     len(elems),
		compare:  cf,
		elements: elems,
	}

	h.heapify()

	return h
}

// heapify makes a heap of the slice in-place
func (h *PrioQ[T]) heapify() {
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

// Insert adds a new element to the priority queue
// The time complexity is  O(log(n)), n = # of elements in the priority queue
func (h *PrioQ[T]) Insert(x T) {
	h.elements = append(h.elements, x)
	h.size++

	// Fix the heap
	i := h.size - 1
	for i >= 0 && h.compare(h.elements[parent(i)], h.elements[i]) {
		h.elements[parent(i)], h.elements[i] = h.elements[i], h.elements[parent(i)]
		i = parent(i)
	}
}

// Extract returns the element at the front of the queue
// It returns an errors whenever the queue is empty
// The time complexity is  O(log(n)), n = # of elements in the queue
func (h *PrioQ[T]) Extract() (T, error) {
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

// IsEmpty indicates whether the queue is empty
func (h *PrioQ[T]) IsEmpty() bool {
	return h.size == 0
}

// Len returns the current size of the priority queue
func (h *PrioQ[T]) Len() int {
    return h.size
}

func (h *PrioQ[T]) largerChild(i int) int {
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
