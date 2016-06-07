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

	f, _ := os.Open("./data/d81")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var grid [80][80]int

	for i := 0; i < 80; i++ {
		for j := 0; j < 80; j++ {
			scanner.Scan()
			text := scanner.Text()
			grid[i][j], _ = strconv.Atoi(text)
		}
	}

	for i := 78; i > -1; i-- {
		grid[79][i] += grid[79][i+1]
		grid[i][79] += grid[i+1][79]
	}

	for i := 78; i > -1; i-- {
		for j := 78; j > -1; j-- {
			if grid[i+1][j] < grid[i][j+1] {
				grid[i][j] += grid[i+1][j]
			} else {
				grid[i][j] += grid[i][j+1]
			}
		}
	}

	fmt.Println(grid[0][0])

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
