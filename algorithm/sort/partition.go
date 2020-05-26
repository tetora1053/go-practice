// n個の整数からなる配列sを、末尾の要素xを基準としてpartitionせよ
package main

import (
	"fmt"
)

func partition(s []int, j int) {
	x := s[len(s)-1]
	i := j - 1
	for j < len(s)-1 {
		if s[j] <= x {
			i++
			s[i], s[j] = s[j], s[i]
		}
		j++
	}
	s[i+1], s[len(s)-1] = s[len(s)-1], s[i+1]
}

func main() {
	var n, num int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		s[i] = num
	}
	partition(s, 0)
	fmt.Println(s)
}
