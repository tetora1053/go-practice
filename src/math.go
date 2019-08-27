package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	max := math.Max(math.Max(a, b), c)

	fmt.Println(max)
}
