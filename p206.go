package main

import (
	"fmt"
	"math/big"
	"time"
)

func IsValid(x []rune) bool {
	for i := 0; i < 9; i++ {
		if x[i*2] != rune('1')+rune(i) {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	ansThree := big.NewInt(0)
	ansThree.SetString("101010103", 10)
	ansSeven := big.NewInt(0)
	ansSeven.SetString("101010107", 10)

	for {
		square := big.NewInt(0)
		square.Set(ansThree)
		square.Mul(square, ansThree)

		if IsValid([]rune(square.String())) {
			fmt.Print(ansThree.String())
			fmt.Println("0")
			break
		} else {
			ansThree.Add(ansThree, big.NewInt(10))
		}

		square.Set(ansSeven)
		square.Mul(square, ansSeven)

		if IsValid([]rune(square.String())) {
			fmt.Print(ansSeven.String())
			fmt.Println("0")
			break
		} else {
			ansSeven.Add(ansSeven, big.NewInt(10))
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
