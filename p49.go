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

func DigitsAreSame(x, y int) bool {
	digits := make([]int, 10)

	for i := 0; i < 10; i++ {
		digits[i] = 0
	}

	for x > 0 {
		digits[x%10]++
		x /= 10
	}

	for y > 0 {
		digits[y%10]--
		y /= 10
	}

	for i := 0; i < 10; i++ {
		if digits[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13

	for maxP < 10000 {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
	}

	for _, i := range prime {
		if i > 999 {
			for _, j := range prime {
				if i < j && DigitsAreSame(i, j) {
					for _, k := range prime {
						if j < k && DigitsAreSame(j, k) && j-i == k-j {
							fmt.Println(i, j, k)
						}
					}
				}
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
