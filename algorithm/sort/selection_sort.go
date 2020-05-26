// 1行目に数列の長さを表す整数Nを、2行目にN個の整数が空白区切りで与えられる。
// 選択ソートのアルゴリズムで昇順に並べ換えよ。
package main

import (
	"fmt"
)

func selectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		minj := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minj] {
				minj = j
			}
		}
		s[i], s[minj] = s[minj], s[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	selectionSort(s)
	fmt.Println(s)
}
