package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	num := int64(28433)

	for i := 0; i < 7830457; i++ {
		num = (num * 2) % 10000000000
	}

	fmt.Println(num + 1)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
