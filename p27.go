package main

import (
	"fmt"
	"math"
	"time"
)

func checkPrime(prime []int, x int) bool {
	if x < 2 {
		return false
	}

	for _, p := range prime {
		if p > int(math.Sqrt(float64(x))) {
			return true
		} else if x%p == 0 {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13
	for maxP < 1000 {
		maxP += 2
		if checkPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
	}

	listB := make([]int, len(prime)*2)
	for i, p := range prime {
		listB[2*i] = p
		listB[2*i+1] = p * -1
	}

	maxMul, maxN := 0, 0

	for _, b := range listB {
		for a := -999; a < 1000; a += 2 {
			for maxP < int(math.Sqrt(float64(maxN*(maxN+a)+b))) {
				maxP += 2
				if checkPrime(prime, maxP) {
					prime = append(prime, maxP)
				}
			}

			if !checkPrime(prime, maxN*(maxN+a)+b) {
				continue
			}

			isPrime := true
			n := 1
			for ; isPrime; n++ {
				curr := n*(n+a) + b

				for maxP < int(math.Sqrt(float64(curr))) {
					maxP += 2
					if checkPrime(prime, maxP) {
						prime = append(prime, maxP)
					}
				}

				isPrime = checkPrime(prime, curr)
			}

			if n > maxN {
				maxMul = a * b
				maxN = n
				fmt.Println(a, b, n)
			}
		}
	}

	fmt.Println(maxMul)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
