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

	num := 1
	den := 1

	for i := 10; i < 100; i++ {
		if i%10 == i/10 || i%10 == 0 {
			continue
		}
		for j := i + 1; j < 100; j++ {
			if j%10 == j/10 || j%10 == 0 {
				continue
			}
			if i%10 == j%10 {
				if i*(j/10) == j*(i/10) {
					fmt.Println(i, "/", j)
					num *= i / 10
					den *= j / 10
				}
			} else if i/10 == j%10 {
				if i*(j/10) == j*(i%10) {
					fmt.Println(i, "/", j)
					num *= i % 10
					den *= j / 10
				}
			} else if i%10 == j/10 {
				if i*(j%10) == j*(i/10) {
					fmt.Println(i, "/", j)
					num *= i / 10
					den *= j % 10
				}
			} else if i/10 == j/10 {
				if i*(j%10) == j*(i%10) {
					fmt.Println(i, "/", j)
					num *= i % 10
					den *= j % 10
				}
			}
		}
	}

	ans := den / GCD(num, den)
	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
