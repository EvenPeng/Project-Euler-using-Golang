package main

import (
	"fmt"
	"time"
)

func SelfPower(base int64) int64 {
	x := base

	for i := int64(1); x != 0 && i < base; i++ {
		x = (x * base) % 10000000000
	}

	return x
}

func main() {
	start := time.Now()

	sum := int64(1)

	for i := int64(2); i < 1001; i++ {
		sum = (sum + SelfPower(i)) % 10000000000
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
