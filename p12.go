package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	prime := []int64{2, 3, 5, 7}
	curr := int64(7)
	sum := int64(28)
	count := int64(1)

	for count < 501 {
		curr++
		sum += curr
		count = 1

		isPrime := true
		for _, p := range prime {
			if curr%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, curr)
		}

		tmp_sum := sum
		for _, p := range prime {
			if tmp_sum%p == 0 {
				tmp_count := int64(2)
				tmp_sum /= p
				for tmp_sum%p == 0 {
					tmp_count++
					tmp_sum /= p
				}
				count *= tmp_count
			}
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
