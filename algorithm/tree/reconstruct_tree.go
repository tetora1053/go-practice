// 螺旋本p203
// 先行順巡回と中間順巡回から木を復元して後行順巡回
package main

import (
	"fmt"
)

var (
	pos int
)

func reconstruct(pre, in []int, post *[]int, l, r int) {
	if l >= r {
		return
	}
	root := pre[pos]
	pos++
	m := 0
	for i := l; i < r; i++ {
		if in[i] == root {
			m = i
			break
		}
	}
	reconstruct(pre, in, post, l, m)
	reconstruct(pre, in, post, m+1, r)
	*post = append(*post, root)
}

func main() {
	var n, num int
	fmt.Scan(&n)

	pre := make([]int, n)
	in := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		pre[i] = num
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		in[i] = num
	}

	post := []int{}
	reconstruct(pre, in, &post, 0, n)

	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(post[i])
	}
	fmt.Print("\n")
}
