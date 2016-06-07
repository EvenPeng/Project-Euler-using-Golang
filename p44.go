package main

import (
	"fmt"
	"math"
	"time"
)

func IsPentagon(x int) bool {
	n := (1 + int(math.Sqrt(float64(1+24*x)))) / 6

	return x == n*(3*n-1)/2
}

func main() {
	start := time.Now()

	min := 0
	first := true

	for i := 2; first; i++ {
		pi := i * (3*i - 1) / 2
		for j := i - 1; j > 0; j-- {
			pj := j * (3*j - 1) / 2
			if IsPentagon(pi+pj) && IsPentagon(pi-pj) {
				min = pi - pj
				first = false
			}
		}
	}

	fmt.Println(min)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
