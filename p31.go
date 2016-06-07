package main

import (
	"fmt"
	"time"
)

func Comb(coins []int, goal, level int) int {
	sum := 0

	for goal >= 0 {
		if level > 1 && goal > 0 {
			sum += Comb(coins, goal, level-1)
		} else {
			sum++
		}

		goal -= coins[level]
	}

	return sum
}

func main() {
	start := time.Now()

	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	goal := 200

	fmt.Println(Comb(coins, goal, len(coins)-1))

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
