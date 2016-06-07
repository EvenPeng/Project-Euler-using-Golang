package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	prev := big.NewInt(1)
	curr := big.NewInt(2)
	next := big.NewInt(3)
	index := 3

	for len(curr.String()) < 1000 {
		next.Add(prev, curr)
		prev.Set(curr)
		curr.Set(next)
		index++
	}

	fmt.Println(index)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
