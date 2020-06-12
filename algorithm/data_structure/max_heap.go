// 最大ヒープ
package main

import (
	"fmt"
)

// 子を持つノードの中で最大のインデックス(h/2)を持つものから降順でmaxHeapifyに渡す
func buildMaxHeap(heap []int) {
	h := len(heap) - 1
	for i := h / 2; i > 0; i-- {
		maxHeapify(heap, i)
	}
}

// iを根とする部分木を最大ヒープにする
func maxHeapify(heap []int, i int) {
	h := len(heap) - 1
	l := i * 2
	r := i*2 + 1
	largest := i
	if l <= h && heap[l] > heap[i] {
		largest = l
	}
	if r <= h && heap[r] > heap[largest] {
		largest = r
	}

	if largest != i {
		heap[largest], heap[i] = heap[i], heap[largest]
		maxHeapify(heap, largest)
	}
}

func main() {
	var h, num int
	fmt.Scan(&h)

	heap := make([]int, h+1)
	for i := 1; i < h+1; i++ {
		fmt.Scan(&num)
		heap[i] = num
	}

	buildMaxHeap(heap)

	for i := 1; i < h+1; i++ {
		fmt.Printf(" %d", heap[i])
	}
	fmt.Print("\n")
}
