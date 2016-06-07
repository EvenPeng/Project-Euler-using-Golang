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

	prime := []int{2, 3, 5, 7, 11, 13, 17}
	digits := make([]int, 10)
	sum := 0

	// init
	for i := 0; i < 9; i++ {
		digits[i] = 9 - i
	}
	digits[9] = 0

	for digits[0] != 0 || digits[9] != 9 {
		satisfy := true
		for i := 0; i < 7; i++ {
			if Dtoi(digits[i+1:], 3)%prime[i] != 0 {
				satisfy = false
				break
			}
		}
		if satisfy {
			sum += Dtoi(digits, 10)
		}

		for i := 9; i > 0; i-- {
			if digits[i] < digits[i-1] {
				for j := 9; j > i-1; j-- {
					if digits[j] < digits[i-1] {
						digits[j], digits[i-1] = digits[i-1], digits[j]
						break
					}
				}
				// reorder
				sort.Sort(sort.Reverse(sort.IntSlice(digits[i:])))
				break
			}
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
