package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d89")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	BtoI := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	//ItoB := map[int]byte{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}

	ans := 0

	for i := 0; i < 1000; i++ {
		scanner.Scan()
		numStr := scanner.Text()
		minStr := []byte{}
		val := BtoI[numStr[len(numStr)-1]]

		for i := 0; i < len(numStr)-1; i++ {
			if BtoI[numStr[i]] < BtoI[numStr[i+1]] {
				val -= BtoI[numStr[i]]
			} else {
				val += BtoI[numStr[i]]
			}
		}

		for val >= 1000 {
			minStr = append(minStr, 'M')
			val -= 1000
		}

		for val >= 100 {
			if val/100 == 9 {
				minStr = append(minStr, 'C', 'M')
				val -= 900
			} else if val/100 >= 5 {
				minStr = append(minStr, 'D')
				val -= 500
			} else if val/100 == 4 {
				minStr = append(minStr, 'C', 'D')
				val -= 400
			} else {
				minStr = append(minStr, 'C')
				val -= 100
			}
		}

		for val >= 10 {
			if val/10 == 9 {
				minStr = append(minStr, 'X', 'C')
				val -= 90
			} else if val/10 >= 5 {
				minStr = append(minStr, 'L')
				val -= 50
			} else if val/10 == 4 {
				minStr = append(minStr, 'X', 'L')
				val -= 40
			} else {
				minStr = append(minStr, 'X')
				val -= 10
			}
		}

		for val > 0 {
			if val == 9 {
				minStr = append(minStr, 'I', 'X')
				val -= 9
			} else if val >= 5 {
				minStr = append(minStr, 'V')
				val -= 5
			} else if val == 4 {
				minStr = append(minStr, 'I', 'V')
				val -= 4
			} else {
				minStr = append(minStr, 'I')
				val -= 1
			}
		}

		ans += len(numStr) - len(minStr)
	}

	fmt.Println(ans)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
