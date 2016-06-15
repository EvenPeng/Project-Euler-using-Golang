package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	curr := big.NewInt(2)
	cnt := 1

	for ; len(curr.String()) < 2; curr.Add(curr, big.NewInt(1)) {
		tmpCurr := big.NewInt(1)
		tmpCurr.Set(curr)

		for exp := 1; len(tmpCurr.String()) == exp; exp++ {
			cnt++
			tmpCurr.Mul(tmpCurr, curr)
		}
	}

	fmt.Println(cnt)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
