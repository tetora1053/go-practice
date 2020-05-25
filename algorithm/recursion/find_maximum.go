// n個の整数からなる配列sの中の最大の要素を出力せよ
package main

import (
	"fmt"
)

func findMax(s []int, l, r int) int {
	if l+1 == r {
		return s[l]
	}
	a := findMax(s, l, (l+r)/2)
	b := findMax(s, (l+r)/2, r)
	return max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var n, num int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		s[i] = num
	}
	fmt.Println(findMax(s, 0, len(s)))
}
