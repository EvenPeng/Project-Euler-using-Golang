package main

import (
	"fmt"
	"time"
)

func Trail(dp []bool, index, curr, goal, currBest int) int {
	best := currBest
	if curr+1 == currBest {
		return currBest
	}

	for i, dpI := range dp {
		if dpI {
			next := curr + 1
			if i+index < goal {
				dp[i+index] = true
				next = Trail(dp, i+index, curr+1, goal, best)
				dp[i+index] = false
			} else if i+index > goal {
				next = goal
			}

			if best > next {
				best = next
			}
		}
	}

	return best
}

func main() {
	start := time.Now()

	ans := 1
	limit := 200
	dp := make([]bool, limit+1)

	for i := 3; i <= limit; i++ {
		for j := 0; j <= i; j++ {
			dp[j] = false
		}
		dp[1] = true
		dp[2] = true

		currStep := Trail(dp, 2, 1, i, i+1)
		ans += currStep
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
