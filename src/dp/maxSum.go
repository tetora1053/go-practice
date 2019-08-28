package main

/*
n個の整数 a[0], a[1], …, a[n−1] が与えられる。これらの整数から何個かの整数を選んで総和をとったときの総和の最大値を求めよ。また、何も選ばない場合の総和は0であるものとする。

1 <= n <= 10000
-1000 <= a[i] <= 1000
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		dp[i+1] = int(math.Max(float64(dp[i]), float64(dp[i]+a)))
	}
	fmt.Println(dp[n])
}
