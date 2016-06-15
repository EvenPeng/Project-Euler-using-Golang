package main

import (
	"fmt"
	"time"
)

func Next(square []int, x int) int {
	sum := 0

	for x > 0 {
		sum += square[x%10]
		x /= 10
	}

	return sum
}

func main() {
	start := time.Now()

	check := make([]bool, 64)
	square := make([]int, 10)
	for i := 0; i < 10; i++ {
		square[i] = i * i
	}

	ans := 0

	check[0] = false
	check[1] = false
	for i := 2; i < 64; i++ {
		curr := i
		for curr != 1 && curr != 89 {
			curr = Next(square, curr)
		}
		if curr == 89 {
			ans++
		}
		check[i] = curr == 89
	}

	for i := 64; i < 10000000; i++ {
		curr := i
		for curr > 63 {
			curr = Next(square, curr)
		}

		if check[curr] {
			ans++
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
