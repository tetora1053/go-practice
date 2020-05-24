package main

import (
	"fmt"
)

func linearSearch(s []int, key int) bool {
	s = append(s, key)
	i := 0
	for s[i] != key {
		i++
	}
	return i != len(s)-1
}

func main() {
	var n, n2, num, key, sum int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		s[i] = num
	}

	fmt.Scan(&n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&key)
		if linearSearch(s, key) {
			sum++
		}
	}
	fmt.Println(sum)
}
