package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(prime []int, x int) bool {
	upperBound := int(math.Sqrt(float64(x)))

	for _, p := range prime {
		if p > upperBound {
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

	prime, maxP := []int{2, 3, 5, 7, 9, 13}, 13

	numP, numN := float64(8), float64(13)
	curr, step := 49, 6
	for numP/numN > 0.1 {
		step += 2
		for maxP < int(math.Sqrt(float64(curr+step*4))) {
			maxP += 2
			if IsPrime(prime, maxP) {
				prime = append(prime, maxP)
			}
		}

		for i := 0; i < 3; i++ {
			curr += step
			if IsPrime(prime, curr) {
				numP++
			}
		}
		curr += step
		numN += 4
		fmt.Println(curr, maxP, numP, numN, numP/numN, step)
	}

	fmt.Println(step + 1)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
