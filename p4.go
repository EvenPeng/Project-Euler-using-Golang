package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	var ans int

	for i := 999; i > 99; i-- {
		ans = i*1000 + i/100 + (i/10%10)*10 + (i%10)*100
		check := false

		for j := int(math.Sqrt(float64(ans))); j > 99; j-- {
			if ans%j == 0 && ans/j < 1000 && ans/j > 99 {
				check = true
				break
			}
		}

		if check {
			break
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
