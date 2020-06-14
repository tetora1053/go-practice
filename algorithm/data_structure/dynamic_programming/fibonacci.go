// フィボナッチ数列のn番目の値を求める(0番目:1, 1番目:1)
// メモ化再帰
package main

import (
	"fmt"
)

var f []int

func fib(n int) int {
	if f[n] != 0 {
		return f[n]
	}
	if n == 0 || n == 1 {
		f[n] = 1
		return 1
	}
	f[n] = fib(n-2) + fib(n-1)
	return f[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	f = make([]int, n+1)
	fmt.Println(fib(n))
}
