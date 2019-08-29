package main

/*
N個の足場があって、i番目の足場の高さはh[i]です。
最初、足場1にカエルがいて、ぴょんぴょん跳ねながら足場Nへと向かいます。カエルは足場iにいるときに
・足場iから足場i+1へと移動する (そのコストは|h[i]−h[i+1]|)
・足場iから足場i+2へと移動する (そのコストは|h[i]−h[i+2]|)
・...
・足場iから足場i+Kへと移動する (そのコストは|h[i]−h[i+k]|)
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
	var N, K int
	fmt.Scanln(&N, &K)

	sc := bufio.NewScanner(os.Stdin)
	heights := make([]int, N+K)
	for i := 0; i < N; i++ {
		sc.Scan()
		heights[i], _ = strconv.Atoi(sc.Text())
	}
	dp := make([]int, N+K)
	for i := range dp {
		dp[i] = 1 << 31
	}

	dp[0] = 0
	for i := 0; i < N; i++ {
		for j := 1; j <= K; j++ {
			dp[i+j] = min(dp[i+j], dp[i] + abs(heights[i+j] - heights[i]))
		}
	}
	fmt.Println(dp[N-1])
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
