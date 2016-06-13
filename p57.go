package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	num := big.NewInt(1)
	den := big.NewInt(2)
	ans := 0

	for i := 1; i < 1000; i++ {
		nextNum := big.NewInt(0).Set(den)
		nextDen := big.NewInt(0).Set(den.Add(den.Add(den, den), num))

		num, den = nextNum, nextDen
		tmpNum := big.NewInt(0).Set(num)
		tmpNum.Add(tmpNum, den)

		if len(tmpNum.String()) > len(den.String()) {
			ans++
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
