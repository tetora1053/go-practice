package main

import (
	"fmt"
)

var cnt int

func mergeSort(s []int, l, r int) {
	mid := (l + r) / 2
	if l+1 < r {
		mergeSort(s, l, mid)
		mergeSort(s, mid, r)
		merge(s, l, mid, r)
	}
}

func merge(s []int, l, mid, r int) {
	n1, n2 := mid-l, r-mid
	s1 := make([]int, n1+1)
	s2 := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		s1[i] = s[l+i]
	}
	s1[n1] = 10000
	for i := 0; i < n2; i++ {
		s2[i] = s[mid+i]
	}
	s2[n2] = 10000

	i, j := 0, 0
	for k := l; k < r; k++ {
		cnt++
		if s1[i] <= s2[j] {
			s[k] = s1[i]
			i++
		} else {
			s[k] = s2[j]
			j++
		}
	}
}

func main() {
	var n, num int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		s[i] = num
	}
	mergeSort(s, 0, len(s))
	fmt.Println(s)
	fmt.Println(cnt)
}
