package main

import (
	"fmt"
	"strconv"
	"time"
)

func IsBouncy(x int64) bool {
	str := strconv.FormatInt(x, 10)

	increase, decrease := true, true

	for i := 1; i < len(str); i++ {
		if str[i-1] < str[i] {
			decrease = false
		}
		if str[i-1] > str[i] {
			increase = false
		}
	}

	return !increase && !decrease
}

func main() {
	start := time.Now()

	curr := int64(100)
	cnt := int64(0)

	for 100*cnt/curr < 99 {
		curr++
		if IsBouncy(curr) {
			cnt++
		}
	}

	fmt.Println(curr)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
