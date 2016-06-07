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

	f, _ := os.Open("./data/d18")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var grid [15][15]int

	for i := 0; i < 15; i++ {
		for j := 0; j <= i; j++ {
			scanner.Scan()
			text := scanner.Text()
			grid[i][j], _ = strconv.Atoi(text)
		}
	}

	for i := 13; i > -1; i-- {
		for j := 0; j <= i; j++ {
			if grid[i+1][j] > grid[i+1][j+1] {
				grid[i][j] += grid[i+1][j]
			} else {
				grid[i][j] += grid[i+1][j+1]
			}
		}
	}

	fmt.Println(grid[0][0])

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
