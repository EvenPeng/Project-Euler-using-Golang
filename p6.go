package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	diff := 5050 * 5050

	for i := 1; i < 101; i++ {
		diff -= i * i
	}

	fmt.Println(diff)

	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}
