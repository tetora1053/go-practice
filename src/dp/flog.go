package main

/*
N個の足場があって、i番目の足場の高さはh[i]です。
最初、足場1にカエルがいて、ぴょんぴょん跳ねながら足場Nへと向かいます。カエルは足場iにいるときに
(1)足場iから足場i+1へと移動する (そのコストは|h[i]−h[i+1]|)
(2)足場iから足場i+2へと移動する (そのコストは|h[i]−h[i+2]|)
のいずれかの行動を選べます。カエルが足場1から足場Nへと移動するのに必要な最小コストを求めよ。

2 <= N <= 10**5
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var N int
	fmt.Scanln(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	h := make([]int, N)
	dp := make([]int, N)

	for i := range h {
		sc.Scan()
		h[i], _ = strconv.Atoi(sc.Text())
	}

	dp[0] = 0
	dp[1] = abs(h[1] - h[0])
	for i := 2; i < N; i++ {
		dp[i] = min(dp[i-2]+abs(h[i]-h[i-2]), dp[i-1]+abs(h[i]-h[i-1]))
	}
	fmt.Println(dp)
	fmt.Println(dp[N-1])
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
