package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main() {
	fmt.Println("test")

	h := make(IntHeap, 0)
	heap.Push(&h, 10)
	heap.Push(&h, 3)
	heap.Push(&h, 9)
	fmt.Println(h)
	heap.Push(&h, 5)
	heap.Push(&h, 1)
	fmt.Println(h)
	heap.Push(&h, 13)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(&h))
	}
}

