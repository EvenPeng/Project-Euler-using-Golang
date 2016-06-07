package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fac := make([]int, 10)
	fac[0] = 1
	for i := 1; i < 10; i++ {
		fac[i] = fac[i-1] * i
	}
	fmt.Println(fac)

	ans := 0

	for curr := 10; curr < fac[9]*7; curr++ {
		sum := 0
		for tmpCurr := curr; tmpCurr > 0; tmpCurr /= 10 {
			sum += fac[tmpCurr%10]
		}

		if sum == curr {
			ans += curr
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
