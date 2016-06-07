package main

import (
	"fmt"
	"sort"
	"time"
)

func Dtoi(digits []int, len int) int {
	sum := 0

	for i := 0; i < len; i++ {
		sum = sum*10 + digits[i]
	}

	return sum
}

func main() {
	start := time.Now()

	digits := make([]int, 9)
	sum := 0

	// init
	for i := 0; i < 9; i++ {
		digits[i] = 9 - i
	}

	for digits[0] != 1 || digits[9-1] != 9 {
		if Dtoi(digits, 1)*Dtoi(digits[1:], 4) == Dtoi(digits[5:], 4) || Dtoi(digits, 2)*Dtoi(digits[2:], 3) == Dtoi(digits[5:], 4) {
			sum += Dtoi(digits[5:], 4)
		}

		for i := 9 - 1; i > 0; i-- {
			if digits[i] < digits[i-1] {
				for j := 9 - 1; j > i-1; j-- {
					if digits[j] < digits[i-1] {
						digits[j], digits[i-1] = digits[i-1], digits[j]
						break
					}
				}
				// reorder
				sort.Sort(sort.Reverse(sort.IntSlice(digits[i:9])))
				break
			}
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
