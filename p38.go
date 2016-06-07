package main

import (
	"fmt"
	"sort"
	"strconv"
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

	// init
	for i := 0; i < 9; i++ {
		digits[i] = 9 - i
	}

	done := false
	for !done {
		for i := 1; !done && i < 5; i++ {
			base := Dtoi(digits[:i], i)
			for j, n := i, 2; j < 9; n++ {
				str := strconv.Itoa(base * n)
				if j+len(str) < 10 {
					if Dtoi(digits[j:j+len(str)], len(str)) == base*n {
						done = true
					} else {
						done = false
						break
					}
					j += len(str)
				} else {
					done = false
					break
				}
			}

			if done {
				fmt.Println(Dtoi(digits, 9), base)
			}
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

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
