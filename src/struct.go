package main

import(
	"fmt"
)

type IntPair struct{
	X, Y int
}

func (int_pair *IntPair) set(x, y int) string {
	int_pair.X = x
	int_pair.Y = y
	return "ok"
}

func main() {
	int_pair := new(IntPair)
	res := int_pair.set(12, 33)
	fmt.Println(res)
	fmt.Println(int_pair)
}