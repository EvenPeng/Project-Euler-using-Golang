package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	sum := 1
	curr := 1

	for i := 2; i < 1001; i += 2 {
		sum += 2 * (curr*2 + i*5)
		curr += i * 4
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
