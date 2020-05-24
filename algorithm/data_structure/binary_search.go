package main

import (
	"fmt"
)

func binarySearch(s []int, key int) bool {
	left, right := 0, len(s)
	for left < right {
		mid := (left + right) / 2
		if s[mid] == key {
			return true
		}
		if key < s[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
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
		if binarySearch(s, key) {
			sum++
		}
	}
	fmt.Println(sum)
}
