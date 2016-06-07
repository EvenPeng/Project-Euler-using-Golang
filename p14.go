package main

import (
	"fmt"
	"time"
)

var count [1000000]int

func Collatz(curr int) int {
	if curr < 1000000 {
		if count[curr] == 0 {
			if curr%2 == 0 {
				count[curr] = Collatz(curr/2) + 1
			} else {
				count[curr] = Collatz(curr+(curr+1)/2) + 2
			}
		}

		return count[curr]
	} else {
		if curr%2 == 0 {
			return Collatz(curr/2) + 1
		} else {
			return Collatz(curr+(curr+1)/2) + 2
		}
	}
}

func main() {
	start := time.Now()

	for i, _ := range count {
		count[i] = 0
	}
	count[1] = 1

	index := 1
	max := 1

	for i := 2; i < 1000000; i++ {
		if Collatz(i) > max {
			max = Collatz(i)
			index = i
		}
	}

	fmt.Println(index)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
