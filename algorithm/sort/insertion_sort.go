// N個の要素を含む数列を昇順に並び替える挿入ソートのプログラムを作成せよ。
// 入力：1行目に数列の長さを表す整数Nが、2行目にN個の整数が空白区切で与えられる。

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	fmt.Println(s)
	for i := 1; i < n; i++ {
		v := s[i]
		j := i - 1
		for j >= 0 && v < s[j] {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = v
		fmt.Println(s)
	}
}
