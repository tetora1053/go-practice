package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("shogi-players.csv")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)

		if err != nil || n == 0 {
			break;
		}

		fmt.Println(string(buf[:n]))
	}
}