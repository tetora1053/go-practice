// 優先度付きキュー
// insert: maxヒープの末尾に要素を追加して親の値と比較、親より値が大きければ交換。この処理を再帰的に行う。
// extract: maxヒープの先頭要素を取り出して、末尾の要素を先頭に移動。maxHeapifyを行う。
package main

import (
	"fmt"
	"math"
)

type pq []int

var (
	h int = 0
)

func (pq pq) insert(key int) {
	h++
	pq[h] = key
	for i := h; i > 1 && pq[i] > pq[i/2]; i = i / 2 {
		pq[i], pq[i/2] = pq[i/2], pq[i]
	}
}

func (pq pq) extract() (ret int) {
	if h == 0 {
		return math.MinInt64
	}
	ret = pq[1]
	pq[1] = pq[h]
	h--
	pq.maxHeapify(1)
	return
}

func (pq pq) maxHeapify(i int) {
	l := i * 2
	r := i*2 + 1
	largest := i
	if l <= h && pq[l] > pq[largest] {
		largest = l
	}
	if r <= h && pq[r] > pq[largest] {
		largest = r
	}

	if i != largest {
		pq[i], pq[largest] = pq[largest], pq[i]
		pq.maxHeapify(largest)
	}
}

func main() {
	var (
		order string
		key   int
	)

	pq := make(pq, 2000000)
	for {
		fmt.Scan(&order)
		switch true {
		case order == "end":
			return
		case order == "insert":
			fmt.Scan(&key)
			pq.insert(key)
		case order == "extract":
			fmt.Println(pq.extract())
		}
	}
}
