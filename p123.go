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

func R(x, y, z int64) int64 {
	r := x
	if r > z {
		r %= z
	}

	for y > 1 {
		r *= x
		y--
		if r > z {
			r %= z
		}
	}

	return r
}

func main() {
	start := time.Now()

	limit := int64(10000000000)
	prime, maxP := []int64{2, 3, 5}, int64(5)

	for maxP*maxP < limit {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
	}

	for {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
			squareP := maxP * maxP
			r := R(maxP-1, int64(len(prime)), squareP)
			r += R(maxP+1, int64(len(prime)), squareP)
			r %= squareP
			if r > limit {
				break
			}
		}
	}

	fmt.Println(len(prime))

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
