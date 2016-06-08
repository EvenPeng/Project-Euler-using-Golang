package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	cnt := 0

	for n := 1; n < 101; n++ {
		comb := n
		for r := 2; r <= n-r; r++ {
			comb *= (n - r + 1)
			comb /= r
			if comb > 1000000 {
				cnt += n - 2*r + 1
				break
			}
		}
	}

	fmt.Println(cnt)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
