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

	limit := 1000000000
	prime, maxP := []int{2, 3, 5}, 5

	for {
		maxP += 2
		if maxP > 100 {
			break
		}
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
	}

	tail := make([]int, len(prime))
	val := make([]int, len(prime))
	list := []int{1}

	for i := 0; i < len(val); i++ {
		val[i] = prime[i]
	}

	for {
		minV := val[0]

		for _, v := range val {
			if minV > v {
				minV = v
			}
		}

		if minV > limit {
			break
		} else {
			list = append(list, minV)
		}

		for i, v := range val {
			if v == minV {
				tail[i]++
				val[i] = prime[i] * list[tail[i]]
			}
		}
	}

	fmt.Println(len(list))

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
