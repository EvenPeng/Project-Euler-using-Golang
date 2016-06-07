package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	ans := 1
	index := 1
	curr := int64(1)
	currStr := strconv.FormatInt(curr, 10)
	pos := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, p := range pos {
		for index < p {
			curr++
			currStr = strconv.FormatInt(curr, 10)
			if index+len(currStr) >= p {
				ans *= int(currStr[p-index-1] - '0')
			}
			index += len(currStr)
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
