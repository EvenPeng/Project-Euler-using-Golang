package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	power := make([]int, 1001)
	for i := 1; i < 1001; i++ {
		power[i] = i * i
	}

	ans := 0
	max := 0

	for p := 5; p < 1001; p++ {
		cnt := 0

		for a := 1; a < p; a++ {
			for b := a; a+b < p; b++ {
				c := (power[p]/2 - a*b) / p
				if c < 1001 && power[a]+power[b] == power[c] {
					cnt++
				}
			}
		}

		if cnt > max {
			max = cnt
			ans = p
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
