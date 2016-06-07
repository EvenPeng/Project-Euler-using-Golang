package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	prime := []int64{3}
	len := 1
	curr := int64(5)

	for len < 10000 {
		isPrime := true
		for _, p := range prime {
			if curr%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			prime = append(prime, curr)
			len++
		}

		curr += 2
	}

	fmt.Println(prime[9999])

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
