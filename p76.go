package main

import (
	"fmt"
	"time"
)

func Comb(x, y int) int {
	sum := 1
	if y > x {
		y = x
	}

	for i := 2; i <= y; i++ {
		sum += Comb(x-i, i)
	}

	return sum
}

func main() {
	start := time.Now()

	fmt.Println(Comb(100, 99))

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
