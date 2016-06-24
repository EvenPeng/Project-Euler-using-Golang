package main

import (
	"fmt"
	"time"
)

func Next(frac []int, curr int) int {
	val := 0

	for curr > 0 {
		val += frac[curr%10]
		curr /= 10
	}

	return val
}

func Chain(frac, dp []int, curr int) int {
	chain := []int{curr}
	next := curr

	for {
		next = Next(frac, next)
		for _, pre := range chain {
			if next == pre {
				dp[curr] = len(chain)
				for i, x := range chain {
					if x < 1000000 {
						dp[x] = dp[curr] - i
					}
				}
				return dp[curr]
			}
		}

		if next < 1000000 && dp[next] != 0 {
			dp[curr] = len(chain) + dp[next]
			for i, x := range chain {
				if x < 1000000 {
					dp[x] = dp[curr] - i
				}
			}
			return dp[curr]
		}

		chain = append(chain, next)
	}

	return dp[curr]
}

func main() {
	start := time.Now()

	frac := make([]int, 10)
	frac[0] = 1
	for i := 1; i < 10; i++ {
		frac[i] = frac[i-1] * i
	}

	dp := make([]int, 1000000)
	dp[0] = 2
	dp[1] = 1
	dp[2] = 1

	ans := 0

	for i := 3; i < 1000000; i++ {
		if dp[i] == 0 {
			if Chain(frac, dp, i) == 60 {
				ans++
			}
		} else if dp[i] == 60 {
			ans++
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
