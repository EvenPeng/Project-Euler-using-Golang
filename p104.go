package main

import (
	"fmt"
	"math/big"
	"time"
)

func IsPandigital(x string) bool {
	check := make([]bool, 10)

	for i := 0; i < 10; i++ {
		check[i] = true
	}

	for _, w := range x {
		if check[int(w-'0')] {
			check[int(w-'0')] = false
		} else {
			return false
		}
	}

	return check[0]
}

func main() {
	start := time.Now()

	fn_1 := big.NewInt(1)
	fn_2 := big.NewInt(1)
	fn_3 := big.NewInt(1)
	k := 2

	for {
		k++
		fn_3.Set(fn_2)
		fn_3.Add(fn_3, fn_1)

		numStr := fn_3.String()
		if len(numStr) > 8 {
			if IsPandigital(numStr[:9]) {
				fmt.Println("front", k)
			}
			if IsPandigital(numStr[len(numStr)-9:]) {
				fmt.Println("tail", k)
			}
			if IsPandigital(numStr[:9]) && IsPandigital(numStr[len(numStr)-9:]) {
				fmt.Println("front and tail", k)
				break
			}
		}

		fn_1.Set(fn_2)
		fn_2.Set(fn_3)
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
