package main

import (
	"fmt"
)

func getClosure() func(int, string) (int, string) {
	var store string

	return func(num int, next string) (int, string) {
		s := store
		store = next
		return num, s
	}
}

func integers() func() int{
	var i int
	return func() int {
		current := i
		i++
		return current
	}
}

func main() {
	f := getClosure()
	fmt.Println(f(1, "Go"))
	fmt.Println(f(2, "Hello"))
	fmt.Println(f(3, "World"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}