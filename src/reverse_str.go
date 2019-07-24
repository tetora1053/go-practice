/*
 reverse chars using recursive function.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]
	reversed_str := reverseStr(str)
	fmt.Println(reversed_str)
}

func reverseStr(str string) string {
	// 1文字ならそのままreturn
	if (len(str) < 2) {
		return str
	}
	// 2文字以上なら先頭１文字を取り除いて再帰
	return reverseStr(str[1:]) + string(str[0])
}