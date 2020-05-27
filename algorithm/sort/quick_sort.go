package main

import (
	"fmt"
)

type c struct {
	suit  string
	value int
}

func partition(cards []c, l, r int) int {
	x := cards[r].value
	j := l - 1
	for i := l; i < r; i++ {
		if cards[i].value <= x {
			j++
			cards[j], cards[i] = cards[i], cards[j]
		}
	}
	cards[j+1], cards[r] = cards[r], cards[j+1]
	return j + 1
}

func quickSort(cards []c, l, r int) {
	if l+1 < r {
		p := partition(cards, l, r-1)
		quickSort(cards, l, p)
		quickSort(cards, p, r)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cards := make([]c, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i].suit, &cards[i].value)
	}
	quickSort(cards, 0, len(cards))
	fmt.Println(cards)
}
