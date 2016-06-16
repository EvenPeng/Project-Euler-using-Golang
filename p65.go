package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	n := 100
	num := big.NewInt(1)
	den := big.NewInt(1)
	if n%3 == 0 {
		den.SetInt64(int64(n / 3 * 2))
	}

	for i := n - 1; i > 1; i-- {
		tmpNum := big.NewInt(0).Set(den)
		if i%3 == 0 {
			tmpNum.Mul(tmpNum, big.NewInt(int64(i/3*2)))
			tmpNum.Add(tmpNum, num)
		} else {
			tmpNum.Add(tmpNum, num)
		}

		num.Set(den)
		den.Set(tmpNum)
	}

	num.Add(num, den)
	num.Add(num, den)
	fmt.Println(num.String(), "/", den.String())

	sum := 0
	numStr := []rune(num.String())

	for _, w := range numStr {
		sum += int(w - '0')
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
