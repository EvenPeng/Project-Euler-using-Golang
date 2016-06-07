package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func IsPrime(prime []int, x int) bool {
	bound := int(math.Sqrt(float64(x)))
	if x%bound == 0 {
		return false
	}

	for _, p := range prime {
		if p >= bound {
			return true
		} else if x%p == 0 {
			return false
		}
	}

	return true
}

func Dtoi(digits []int, len int) int {
	sum := 0

	for i := 0; i < len; i++ {
		sum = sum*10 + digits[i]
	}

	return sum
}

func main() {
	start := time.Now()

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13
	for maxP < int(math.Sqrt(float64(987654321))) {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
	}

	digits := make([]int, 9)
	done := false

	for n := 9; !done && n > 3; n-- {
		// init
		for i := 0; i < n; i++ {
			digits[i] = n - i
		}

		for {
			curr := Dtoi(digits, n)
			if curr%2 != 0 && IsPrime(prime, curr) {
				fmt.Println(curr)
				done = true
				break
			}

			for i := n - 1; i > 0; i-- {
				if digits[i] < digits[i-1] {
					for j := n - 1; j > i-1; j-- {
						if digits[j] < digits[i-1] {
							digits[j], digits[i-1] = digits[i-1], digits[j]
							break
						}
					}
					// reorder
					sort.Sort(sort.Reverse(sort.IntSlice(digits[i:n])))
					break
				}
			}

			if digits[0] == 1 && digits[n-1] == n {
				break
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
