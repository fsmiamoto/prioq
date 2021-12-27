package prioq

import (
	"reflect"
	"testing"
)

func MaxInt(a, b int) bool {
	return a < b
}

func MinInt(a, b int) bool {
	return a > b
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		cf       CompareFunc[int]
		capacity int
		input    []int
		expect   []int
	}{
		{"MaxHeap #1", MaxInt, 7,
			[]int{10, 20, 15, 12, 40, 25, 18},
			[]int{40, 20, 25, 12, 10, 15, 18}},
		{"MaxHeap #2", MaxInt, 20,
			[]int{100, 88, 24, 43, 35, 7, 97, 39, 46, 38, 15, 53, 65, 93, 87, 84, 59, 52, 24, 53},
			[]int{100, 88, 97, 84, 53, 65, 93, 59, 52, 38, 15, 53, 7, 24, 87, 39, 43, 46, 24, 35}},
		{"MaxHeap #3", MaxInt, 2,
			[]int{10, 20},
			[]int{20, 10}},
		{"MaxHeap #4", MaxInt, 1,
			[]int{5},
			[]int{5}},
		{"MaxHeap #5", MaxInt, 10000,
			[]int{},
			[]int{}},
		{"MinHeap #1", MinInt, 12,
			[]int{8, 12, 9, 7, 22, 3, 26, 14, 11, 15, 22},
			[]int{3, 7, 8, 11, 15, 9, 26, 14, 12, 22, 22}},
		{"MinHeap #2", MinInt, 5,
			[]int{40, 30, 50, 100, 15},
			[]int{15, 30, 50, 100, 40}},
		{"MinHeap #3", MinInt, 22,
			[]int{15, 7, 96, 38, 54, 52, 69, 25, 28, 46, 18, 39, 6, 7, 29, 88, 40, 37, 67, 22},
			[]int{6, 7, 7, 25, 18, 39, 15, 38, 28, 22, 54, 96, 52, 69, 29, 88, 40, 37, 67, 46}},
		{"MinHeap #3", MinInt, 5,
			[]int{10, 20},
			[]int{10, 20}},
		{"MinHeap #4", MinInt, 2,
			[]int{5},
			[]int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewWithCompareFunc(tt.input, tt.cf)

			if len(heap.elements) != len(tt.expect) {
				t.Errorf("Heap has the wrong number of elements")
			}

			equal := reflect.DeepEqual(heap.elements, tt.expect)
			if !equal {
				t.Errorf("Expected %v but got %v", tt.expect, heap.elements)
			}
		})
	}

}

func TestExtract(t *testing.T) {
	tests := []struct {
		name             string
		shouldFail       bool
		extractedElement int
		heap             *PrioQ[int]
		expect           *PrioQ[int]
	}{
		{"Extract on MaxHeap", false, 50,
			&PrioQ[int]{
				size:     7,
				compare:  MaxInt,
				elements: []int{50, 30, 20, 15, 10, 8, 16},
			},
			&PrioQ[int]{
				size:     6,
				compare:  MaxInt,
				elements: []int{30, 16, 20, 15, 10, 8, 50},
			},
		},
		{"Extract on small MaxHeap", false, 50,
			&PrioQ[int]{
				size:     2,
				compare:  MaxInt,
				elements: []int{50, 30},
			},
			&PrioQ[int]{
				size:     1,
				compare:  MaxInt,
				elements: []int{30, 50},
			},
		},
		{"Extract on MinHeap", false, 3,
			&PrioQ[int]{
				size:     11,
				compare:  MinInt,
				elements: []int{3, 7, 8, 11, 15, 9, 26, 14, 12, 22, 22},
			},
			&PrioQ[int]{
				size:     10,
				compare:  MinInt,
				elements: []int{7, 11, 8, 12, 15, 9, 26, 14, 22, 22, 3},
			},
		},
		{"Empty heap", true, 0, 
			&PrioQ[int]{
				size:     0,
				compare:  MinInt,
				elements: []int{},
			},
			&PrioQ[int]{
				size:     0,
				compare:  MinInt,
				elements: []int{},
			},
		},
	}

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            element, err := tt.heap.Extract()
            assertEqual(t, element, tt.extractedElement)
            assertError(t, tt.shouldFail, err)
            assertEqualHeap(t, tt.heap, tt.expect)
        })
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name    string
		element int
		heap    *PrioQ[int]
		expect  *PrioQ[int]
	}{
		{"Insert on MaxHeap", 60,
			&PrioQ[int]{
				size:     7,
				compare:  MaxInt,
				elements: []int{50, 30, 20, 15, 10, 8, 16},
			},
			&PrioQ[int]{
				size:     8,
				compare:  MaxInt,
				elements: []int{60, 50, 20, 30, 10, 8, 16, 15},
			},
		},
		{"Insert on MaxHeap #2", 60,
			&PrioQ[int]{
				size:     9,
				compare:  MaxInt,
				elements: []int{50, 30, 20, 15, 10, 8, 16, 9, 8},
			},
			&PrioQ[int]{
				size:     10,
				compare:  MaxInt,
				elements: []int{60, 50, 20, 15, 30, 8, 16, 9, 8, 10},
			},
		},
		{"Insert on MinHeap #1", 1,
			&PrioQ[int]{
				size:     4,
				compare:  MinInt,
				elements: []int{3, 5, 7, 10},
			},
			&PrioQ[int]{
				size:     5,
				compare:  MinInt,
				elements: []int{1, 3, 7, 10, 5},
			},
		},
		{"Insert on MinHeap #2", 1,
			&PrioQ[int]{
				size:     0,
				compare:  MinInt,
				elements: []int{},
			},
			&PrioQ[int]{
				size:     1,
				compare:  MinInt,
				elements: []int{1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Insert(tt.element)
			assertEqualHeap(t, tt.heap, tt.expect)
		})
	}
}

func assertEqual(t *testing.T, got, expect interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Element is not the expected one, got %v but expected %v", got, expect)
	}
}

func assertError(t *testing.T, shouldFail bool, err error) {
	t.Helper()
	if shouldFail && err == nil {
		t.Errorf("A error was expected but got none")
	}

	if !shouldFail && err != nil {
		t.Errorf("No error was expected but got %v", err)
	}
}

func assertEqualHeap(t *testing.T, got, want *PrioQ[int]) {
	t.Helper()
	if got.size != want.size {
		t.Errorf("Heaps have different sizes, got %d want %d", got.size, want.size)
	}
	if !reflect.DeepEqual(got.elements, want.elements) {
		t.Errorf("Heaps have different elements, got %v want %v", got.elements, want.elements)
	}
}
