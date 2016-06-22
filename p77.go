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

func Comb(prime []int, x, y int) int {
	sum := 0

	if y > x {
		y = x
	}

	for _, p := range prime {
		if p > y {
			break
		} else {
			if x-p == 0 {
				sum++
			} else {
				sum += Comb(prime, x-p, p)
			}
		}
	}

	return sum
}

func main() {
	start := time.Now()

	prime, maxP := []int{2, 3, 5}, 5

	for {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
		if Comb(prime, maxP-1, maxP-2) > 5000 {
			fmt.Println(maxP - 1)
			break
		}
		if Comb(prime, maxP, maxP-1) > 5000 {
			fmt.Println(maxP)
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
