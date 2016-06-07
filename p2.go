package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	prev := 1
	curr := 2
	next := prev + curr

	sum := 2

	for next <= 4000000 {
		prev = curr
		curr = next
		next = prev + curr
		if curr%2 == 0 {
			sum += curr
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}
