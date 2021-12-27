package prioq_test

import (
	"fmt"
	"log"

	"github.com/fsmiamoto/prioq"
)

// This example shows a use in the Heapsort algorithm
func ExampleHeapSort() {
	values := []int{40, 30, 50, 100, 15}

	pq := prioq.New[int](values)

	var sorted []int
	for !pq.IsEmpty() {
		v, err := pq.Extract()
		if err != nil {
			log.Fatal(err)
		}
		sorted = append(sorted, v)
	}
	fmt.Println(sorted)
	// Output:
	// [15 30 40 50 100]
}
