# prioq

A priority queue package using a binary heap data structure in Go.

Built with [Go 1.18 Generics](https://go.dev/blog/why-generics) support.

Adapted from [heap](https://github.com/fsmiamoto/heap)

The API is still at an experimental phase.

## Requirements

Go 1.18

If you've already have Go >=1.16i installed, you can get the beta with:
```sh
$ make go1.18
```

## Example
```go
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
```
