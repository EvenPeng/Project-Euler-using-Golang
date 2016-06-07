package main

import (
	"fmt"
	"math"
	"time"
)

func IsSquare(x int) bool {
	root := int(math.Sqrt(float64(x)))

	return root*root == x
}

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

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13

	done := false
	for !done {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		} else {
			satisfy := false
			for i := len(prime) - 1; i > -1; i-- {
				curr := maxP - prime[i]
				if curr%2 == 0 && IsSquare(curr/2) {
					satisfy = true
					break
				}
			}
			if !satisfy {
				done = true
			}
		}
	}

	fmt.Println(maxP)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
