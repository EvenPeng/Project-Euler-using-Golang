package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	limit := int64(1000000)
	phi := make([]int64, limit+1)
	ans := int64(0)

	for i := int64(1); i < limit+1; i++ {
		phi[i] = i
	}

	for i := int64(2); i < limit+1; i++ {
		if phi[i] == i {
			for j := i; j < limit+1; j += i {
				phi[j] = phi[j] * (i - 1) / i
			}
		}
		ans += phi[i]
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
