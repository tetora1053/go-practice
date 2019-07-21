package main

import(
	"fmt"
)

func main() {
	var m interface{} = true
	b := m.(bool)
	fmt.Println(b)

	var n interface{} = 3
	i := n.(int)
	fmt.Println(i)

	var o interface{} = "abc"
	u, isUint := o.(uint)
	fmt.Println(u, isUint)

	switch v := n.(type) {
	case bool:
		fmt.Println("bool:", v)
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	default:
		fmt.Println("unknown")
	}
}
