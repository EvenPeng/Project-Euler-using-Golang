package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	var prime []int64
	for i := int64(3); i <= int64(math.Sqrt(600851475143)); i += 2 {
		if 600851475143%i == 0 {
			isPrime := true
			for _, p := range prime {
				if i%p == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				prime = append(prime, i)
			}
		}
	}

	fmt.Println(prime[len(prime)-1])

	fmt.Println(prime)

	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}
