package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(prime []int, x int) bool {
	bound := int(math.Sqrt(float64(x)))
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

	prime, maxP := []int{2, 3}, 3
	maxN := 6
	done := false

	for !done {
		isPrime := false
		for !isPrime {
			maxP += 2
			if IsPrime(prime, maxP) {
				prime = append(prime, maxP)
				isPrime = true
				if maxN*maxP < 1000001 {
					maxN *= maxP
				} else {
					done = true
				}
			}
		}
	}

	fmt.Println(maxN)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
