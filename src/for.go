package main

import (
	"fmt"
)

func main() {
	fmt.Println("クラシックなfor")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("just for")
	var j int
	for {
		fmt.Println(j)
		j++
		if j >= 10 {
			break
		}
	}

	fmt.Println("条件式付きfor")
	var k int
	for k < 10 {
		fmt.Println(k)
		k++
	}

	fmt.Println("rangeによるfor")
	nums := [...]int{0,1,2,3,4,5,6,7,8,9}
	for _, num := range nums {
		fmt.Println(num)
	}

	fmt.Println("多重のforから一気に抜ける")
	LOOP:
	for l := 0; l < 10; l++ {
		for m := 0; m < 10; m++ {
			if l == 3 {
				break LOOP
			}
			fmt.Printf("%d %d\n", l, m)
		}
	}
}
