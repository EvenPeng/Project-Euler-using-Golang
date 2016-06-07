package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	num := big.NewInt(1)

	for i := int64(2); i < 101; i++ {
		num.Mul(num, big.NewInt(i))
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
