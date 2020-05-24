// 逆ポーランド記法の入力が与えられるので、スタックを使って計算し、計算結果を出力せよ。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	top   int
	stack []int
)

func init() {
	stack = make([]int, 100)
}

func push(x int) {
	top++
	stack[top] = x
}

func pop() int {
	top--
	return stack[top+1]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	s := strings.Split(str, " ")
	for _, v := range s {
		if v == "+" {
			b, a := pop(), pop()
			push(a + b)
		} else if v == "-" {
			b, a := pop(), pop()
			push(a - b)
		} else if v == "*" {
			b, a := pop(), pop()
			push(a * b)
		} else {
			x, _ := strconv.Atoi(v)
			push(x)
		}
	}
	fmt.Println(pop())
}
