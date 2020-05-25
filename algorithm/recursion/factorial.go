// 入力に整数nを受け取り、nの階乗を出力せよ
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}
