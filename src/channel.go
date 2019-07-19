package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	ch <- 5
	ch <- 10
	ch <- 11
	close(ch)
	var (
		i int
		ok bool
	)

	i, ok = <- ch
	fmt.Println("ok", ok)
	fmt.Println(i)

	i, ok = <- ch
	fmt.Println("ok", ok)
	fmt.Println(i)

	i, ok = <- ch
	fmt.Println("ok", ok)
	fmt.Println(i)

	i, ok = <- ch
	fmt.Println("ok", ok)
	fmt.Println(i)
}