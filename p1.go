package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	sum := 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elapsd:", elapsed)
}
