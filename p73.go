package main

import (
	"fmt"
	"time"
)

func GCD(x, y int) int {
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}

	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	start := time.Now()

	D := 12000
	ans := 0

	for d := 4; d <= D; d++ {
		n := d / 2
		if d%2 == 0 {
			n--
		}
		for ; d < 3*n; n-- {
			if GCD(n, d) == 1 {
				ans++
			}
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
