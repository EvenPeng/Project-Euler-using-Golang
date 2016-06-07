package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d82")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	const size = 80
	var grid [size][size]int
	var minRow [size]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			scanner.Scan()
			text := scanner.Text()
			grid[i][j], _ = strconv.Atoi(text)
		}
		minRow[i] = grid[i][size-1]
	}

	for j := size - 2; j > -1; j-- {
		for i := 0; i < size; i++ {
			minRow[i] += grid[i][j]
		}

		for i := 1; i < size; i++ {
			if minRow[i] > minRow[i-1]+grid[i][j] {
				minRow[i] = minRow[i-1] + grid[i][j]
			}
		}
		for i := size - 2; i > -1; i-- {
			if minRow[i] > minRow[i+1]+grid[i][j] {
				minRow[i] = minRow[i+1] + grid[i][j]
			}
		}
	}

	min := minRow[0]

	for i := 1; i < size; i++ {
		if min > minRow[i] {
			min = minRow[i]
		}
	}

	fmt.Println(min)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
