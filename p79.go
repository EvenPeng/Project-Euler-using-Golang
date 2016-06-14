package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func IsValid(x, y []rune) bool {
	indexX, indexY := 0, 0

	for indexX < len(x) && indexY < len(y) {
		if x[indexX] == y[indexY] {
			indexY++
		}
		indexX++
	}

	return indexY == len(y)
}

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d79")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var code [][]rune

	for i := 0; i < 50; i++ {
		scanner.Scan()
		code = append(code, []rune(scanner.Text()))
	}

	var link [10][10]bool

	for i := 0; i < 50; i++ {
		link[code[i][0]-'0'][code[i][1]-'0'] = true
		link[code[i][1]-'0'][code[i][2]-'0'] = true
	}

	var table [10]bool
	var inDeg, outDeg [10]int

	for i := 0; i < 10; i++ {
		table[i] = false
		inDeg[i] = 0
		outDeg[i] = 0
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if link[i][j] {
				table[i] = true
				table[j] = true
				outDeg[i]++
				inDeg[j]++
			}
		}
	}

	left := 0
	for _, x := range table {
		if x {
			left++
		}
	}

	for left > 0 {
		curr := 0
		for curr < 10 {
			if table[curr] && inDeg[curr] == 0 {
				fmt.Print(curr)

				table[curr] = false
				left--

				for i := 0; i < 10; i++ {
					if link[i][curr] {
						outDeg[i]--
						inDeg[curr]--
					}
					if link[curr][i] {
						outDeg[curr]--
						inDeg[i]--
					}
				}

				break
			} else {
				curr++
			}
		}
	}
	fmt.Println()

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
