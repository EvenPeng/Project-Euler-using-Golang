package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	power := make([]int64, 10)
	for i := 0; i < 10; i++ {
		power[i] = int64(math.Pow(float64(i), float64(5)))
	}

	curr := int64(10)
	sum := int64(0)

	for curr <= power[9]*10 {
		curr++
		tmp_sum := int64(0)
		for x := curr; x > 0; x /= 10 {
			tmp_sum += power[x%10]
		}
		if tmp_sum == curr {
			sum += curr
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
