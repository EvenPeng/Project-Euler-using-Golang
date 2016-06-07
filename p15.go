package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var grid [21][21]int

	for i := 0; i < 21; i++ {
		grid[i][0] = 1
		grid[0][i] = 1
	}

	for i := 1; i < 21; i++ {
		for j := 1; j < i; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
		grid[i][i] = grid[i][i-1] * 2
	}

	fmt.Println(grid[20][20])

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
