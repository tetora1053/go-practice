// FX取引。ある通貨について時刻tにおける価格Rtが入力として与えられる。
// 価格の差Rj-Ri(j>i)の最大値を求めよ。
// 入力：最初の行に整数n、続くn行に整数Rtが順番に与えられる。

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, minV, maxProf, cur int
	fmt.Scan(&n)

	maxProf = math.MinInt64
	fmt.Scan(&minV)
	for i:=1; i<n; i++ {
		fmt.Scan(&cur)
		maxProf = max(maxProf, cur - minV)
		minV = min(minV, cur)
	}
	fmt.Println(maxProf)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

