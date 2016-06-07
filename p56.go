package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	max := 1

	for a := int64(2); a < 100; a++ {
		base := big.NewInt(a)
		num := big.NewInt(a)
		for b := 1; b < 100; b++ {
			num.Mul(num, base)
			numStr := num.String()
			sum := 0
			for _, d := range numStr {
				sum += int(d - '0')
			}
			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
