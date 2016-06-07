package main

import (
	"fmt"
	"time"
)

func LCM(x, y int) int {
	for {
		if x == 0 || y == 0 {
			break
		}
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

	ans := 6

	for i := 4; i < 20; i++ {
		lcm := LCM(ans, i)
		ans = ans * i / lcm
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}
