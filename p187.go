package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(prime []int64, x int64) bool {
	bound := int64(math.Sqrt(float64(x)))
	if x%bound == 0 {
		return false
	}

	for _, p := range prime {
		if p > bound {
			return true
		}
		if x%p == 0 {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	limit := int64(100000000)
	prime, maxP, tail := []int64{2, 3}, int64(3), 1
	ans := 3

	for maxP < limit/2 {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)

			if maxP*maxP < limit {
				ans += len(prime)
				tail++
			} else {
				for prime[tail]*maxP >= limit {
					tail--
				}

				ans += tail + 1
			}
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
