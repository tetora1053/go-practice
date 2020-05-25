// 長さnの数列aと整数mに対して、aの要素の中のいくつかの要素を足し合わせてmが作れるかどうか判定せよ。
// 入力：1行目にn、2行目にaを表すn個の整数、3行目にq、4行目にq個の整数mが与えられる。
// 出力：各mについてAの要素を足し合わせてmを作ることができればyes,できなければnoと出力せよ。
package main

import (
	"fmt"
)

func solve(a []int, i int, m int) (res bool) {
	if m == 0 {
		return true
	} else if i == len(a) {
		return false
	}
	res = solve(a, i+1, m) || solve(a, i+1, m-a[i])
	return
}

func main() {
	var n, num, q, m int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		a[i] = num
	}
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&m)
		if solve(a, 0, m) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
