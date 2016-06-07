package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	num := []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4, 3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
	// 20 twenty
	num = append(num, 6)
	for i := 1; i < 10; i++ {
		num = append(num, 6+num[i])
	}
	// 30 thirty
	num = append(num, 6)
	for i := 1; i < 10; i++ {
		num = append(num, 6+num[i])
	}
	// 40 forty
	num = append(num, 5)
	for i := 1; i < 10; i++ {
		num = append(num, 5+num[i])
	}
	// 50 fifty
	num = append(num, 5)
	for i := 1; i < 10; i++ {
		num = append(num, 5+num[i])
	}
	// 60 sixty
	num = append(num, 5)
	for i := 1; i < 10; i++ {
		num = append(num, 5+num[i])
	}
	// 70 seventy
	num = append(num, 7)
	for i := 1; i < 10; i++ {
		num = append(num, 7+num[i])
	}
	// 80 eighty
	num = append(num, 6)
	for i := 1; i < 10; i++ {
		num = append(num, 6+num[i])
	}
	// 90 ninety
	num = append(num, 6)
	for i := 1; i < 10; i++ {
		num = append(num, 6+num[i])
	}

	sum := 0
	for i := 1; i < 100; i++ {
		sum += num[i]
	}

	for i := 1; i < 10; i++ {
		sum_hunder += 7 + num[i]
	}

	sum = sum*10 + sum_hunder*100 + 3*99*9 + 11

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
