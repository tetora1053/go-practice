package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// 引数にファイル名指定
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	/*
	 * Read
	 */
	buf := make([]byte, 1024)
	// 全部読み込み
	n, err := f.Read(buf)

	if err != nil || n == 0 {
		panic(err);
	}

	row := string(buf[:n])
	rowSlice := strings.Split(row, ",")
	fmt.Println(rowSlice)

	fmt.Println("-------------------")

	/*
	 * Scan
	 */
	 // Seek()でoffsetを最初に戻す
	 f.Seek(0, 0)
	 scanner := bufio.NewScanner(f)
	 for i := 0; scanner.Scan();{
		i++
		// ヘッダーはとばす
		if i == 1 {
			continue
		}
		// 1行ずつ
		rowSlice := strings.Split(scanner.Text(), ",")
		fmt.Println(rowSlice)
	 }
}

