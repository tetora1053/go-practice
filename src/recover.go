package main

import(
	"fmt"
)

func recoverPanic() {
	res := recover()
	if (res != nil) {
		fmt.Println(res)
	}
}

func main() {
	defer recoverPanic()

	panic("panic!")

	fmt.Println("not through here")
}
