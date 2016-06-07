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

	f, _ := os.Open("./data/d11")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	max := 0
	var grid [20][20]int

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			scanner.Scan()
			text := scanner.Text()
			grid[i][j], _ = strconv.Atoi(text)
		}
	}

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			// vertical
			curr := grid[i][j]
			for k := 1; k < 4; k++ {
				curr *= grid[i][j+k]
			}
			if curr > max {
				max = curr
			}
			// diagonal \
			curr = grid[i][j]
			for k := 1; k < 4; k++ {
				curr *= grid[i+k][j+k]
			}
			if curr > max {
				max = curr
			}
			// horizontal
			curr = grid[i][j]
			for k := 1; k < 4; k++ {
				curr *= grid[i+k][j]
			}
			if curr > max {
				max = curr
			}
		}
	}

	for i := 3; i < 20; i++ {
		for j := 0; j < 16; j++ {
			// diagonal /
			curr := grid[i][j]
			for k := 1; k < 4; k++ {
				curr *= grid[i-k][j+k]
			}
			if curr > max {
				max = curr
			}
		}
	}

	fmt.Println(max)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
