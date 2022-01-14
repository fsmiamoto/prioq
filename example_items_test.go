package prioq_test

import (
	"fmt"
	"log"

	"github.com/fsmiamoto/prioq"
)

// Item is an example defined type
type Item struct {
	key int
}

// MaxItem is an example CompareFunc that builds a MaxHeap using the key property
func MaxItem(node, child Item) bool {
	return child.key > node.key
}

// This example shows how to implement your own CompareFunc and then use it on
// a defined type
func Example_items() {
	values := []Item{
		Item{key: 8},
		Item{key: 22},
		Item{key: 3},
		Item{key: 14},
		Item{key: 42},
	}

	pq := prioq.NewWithCompareFunc[Item](values, MaxItem)

	for !pq.IsEmpty() {
		item, err := pq.Extract()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(item.key)
	}
	// Output:
	// 42
	// 22
	// 14
	// 8
	// 3
}
