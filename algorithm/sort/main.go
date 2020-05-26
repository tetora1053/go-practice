package main

import (
	"fmt"
	"strconv"
)

type C struct {
	suit string
	val  int
}

func bubble(s []C) {
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[j].val < s[j-1].val {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

func selection(s []C) {
	for i := 0; i < len(s); i++ {
		minj := i
		for j := i + 1; j < len(s); j++ {
			if s[j].val < s[minj].val {
				minj = j
			}
		}
		s[i], s[minj] = s[minj], s[i]
	}
}

func isStable(s1, s2 []C) {
	for i := 0; i < len(s1); i++ {
		if s1[i].suit != s2[i].suit {
			fmt.Println("Not Stable")
			return
		}
	}
	fmt.Println("Stable")
}

func main() {
	var (
		n   int
		str string
	)
	fmt.Scan(&n)
	s1 := make([]C, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		suit := string(str[0])
		val, _ := strconv.Atoi(string(str[1]))
		s1[i].suit = suit
		s1[i].val = val
	}
	s2 := make([]C, n)
	copy(s2, s1)

	bubble(s1)
	fmt.Println(s1)
	fmt.Println("Stable")
	selection(s2)
	fmt.Println(s2)
	isStable(s1, s2)
}
