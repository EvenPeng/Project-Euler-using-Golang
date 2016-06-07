package main

import (
	"fmt"
	"math"
	"time"
)

func isAbundant(x int) bool {
	sum := 1

	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			sum += i
			if x != i*i {
				sum += x / i
			}
		}
	}

	return sum > x
}

func main() {
	start := time.Now()

	abundant := []int{12}

	for i := 13; i < 28112; i++ {
		if isAbundant(i) {
			abundant = append(abundant, i)
		}
	}

	check := make([]bool, 28124)
	for i := 0; i < 28124; i++ {
		check[i] = true
	}

	for i := 0; i < len(abundant); i++ {
		for j := i; j < len(abundant) && abundant[i]+abundant[j] < 28124; j++ {
			check[abundant[i]+abundant[j]] = false
		}
	}

	sum := 0

	for i, yes := range check {
		if yes {
			sum += i
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
