package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	check := make(map[int64]bool)
	limit := int64(1000000000000)
	bound := (int64(math.Sqrt(float64(4*limit-1)) - 1)) / 2

	for curr := int64(2); curr <= bound; curr++ {
		for val := (curr+1)*curr + 1; val < limit; val = val*curr + 1 {
			check[val] = true
		}
	}

	sum := int64(1)

	for i, _ := range check {
		sum += i
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
