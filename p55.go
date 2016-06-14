package main

import (
	"fmt"
	"strconv"
	"time"
)

func runeReverse(x []rune) []rune {
	y := make([]rune, len(x))

	for i := 0; i < len(x); i++ {
		y[i] = x[len(x)-1-i]
	}

	return y
}

func runeAdd(x, y []rune) []rune {
	z := make([]rune, len(x)+1)

	for i := len(x) - 1; i > -1; i-- {
		z[i+1] = x[i] + y[i] - '0'
	}

	z[0] = 0
	for i := len(z) - 1; i > 0; i-- {
		if z[i] > '9' {
			z[i] -= 10
			z[i-1]++
		}
	}

	if z[0] == 0 {
		return z[1:]
	} else {
		z[0] += '0'
		return z
	}
}

func IsPalindrome(x []rune) bool {
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-1-i] {
			return false
		}
	}

	return true
}

func IsLychrel(x []rune) bool {
	for i := 0; i < 50; i++ {
		x = runeAdd(x, runeReverse(x))
		if IsPalindrome(x) {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	ans := 0

	for i := int64(10); i < 10000; i++ {
		if IsLychrel([]rune(strconv.FormatInt(i, 10))) {
			ans++
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
