package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	var num, base big.Int
	num.Set(big.NewInt(1))
	base.Set(big.NewInt(2))

	for i := 0; i < 1000; i++ {
		num.Mul(&num, &base)
	}

	num_str := num.String()

	sum := 0

	for _, d := range num_str {
		sum += int(d - '0')
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
