// バイナリーヒープ、完全二分木
// 1オリジンの配列
// 親のインデックス：現在のインデックス / 2
// 左の子：現在のインデックス * 2
// 右の子：現在のインデックス * 2 + 1
package main

import (
	"fmt"
)

func parent(id int) int {
	return id / 2
}

func left(id int) int {
	return id * 2
}

func right(id int) int {
	return id*2 + 1
}

func main() {
	var h, num int
	fmt.Scan(&h)

	bh := make([]int, h+1)
	for i := 1; i < len(bh); i++ {
		fmt.Scan(&num)
		bh[i] = num
	}

	for i := 1; i < len(bh); i++ {
		fmt.Printf("node %d: key = %d, ", i, bh[i])
		if parent(i) > 0 {
			fmt.Printf("parent key = %d, ", bh[parent(i)])
		}
		if left(i) < len(bh) {
			fmt.Printf("left key = %d, ", bh[left(i)])
		}
		if right(i) < len(bh) {
			fmt.Printf("right key = %d, ", bh[right(i)])
		}
		fmt.Print("\n")
	}
}
