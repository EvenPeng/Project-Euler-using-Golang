package main

import (
	"fmt"
	"math"
	"time"
)

func d(x int) int {
	sum := 1

	for p := 2; p <= int(math.Sqrt(float64(x))); p++ {
		if x%p == 0 {
			if x != p*p {
				sum += p + x/p
			} else {
				sum += p
			}
		}
	}

	return sum
}

func main() {
	start := time.Now()

	check := make([]bool, 10000)

	for i := 1; i < 10000; i++ {
		check[i] = false
	}

	sum := 0
	for i := 2; i < 10000; i++ {
		if !check[i] {
			d_a := d(i)
			if d_a <= i {
				continue
			}
			d_b := d(d_a)
			if i == d_b {
				sum += d_a + d_b
				check[d_a] = true
				check[d_b] = true
			}
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
