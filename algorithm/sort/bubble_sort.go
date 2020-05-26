package main

import (
	"fmt"
)

func bubbleSort(s []int) (sw int) {
	flag := true
	i := 0
	for flag {
		flag = false
		for j := len(s) - 1; j > i; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				flag = true
				sw++
			}
		}
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	sw := bubbleSort(s)
	fmt.Println(s)
	fmt.Println(sw)
}
