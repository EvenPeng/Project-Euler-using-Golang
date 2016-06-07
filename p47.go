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

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13
	curr := 210
	cnt := 0

	for cnt < 4 {
		curr++
		tmpCurr := curr
		num := 0

		for maxP < int(math.Sqrt(float64(curr))) {
			maxP += 2
			if IsPrime(prime, maxP) {
				prime = append(prime, maxP)
			}
		}

		for _, p := range prime {
			if p*p > curr {
				break
			}
			if tmpCurr%p == 0 {
				num++
				for tmpCurr%p == 0 {
					tmpCurr /= p
				}
			}
		}

		if num == 4 || (num == 3 && tmpCurr > 1) {
			cnt++
		} else {
			cnt = 0
		}
	}

	fmt.Println(curr - 3)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
