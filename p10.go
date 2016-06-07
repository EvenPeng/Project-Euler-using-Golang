package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	check := make([]bool, 2000000)
	for i, _ := range check {
		check[i] = true
	}

	sum := int64(2)

	for curr := int64(3); curr < 2000000; curr += 2 {
		if check[curr] {
			sum += curr
			for mark := curr; mark < 2000000; mark += curr {
				check[mark] = false
			}
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
