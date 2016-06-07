package main

import (
	"fmt"
	"time"
)

func isLeap(x int) bool {
	if x%400 == 0 {
		return true
	} else if x%100 == 0 {
		return false
	} else {
		return x%4 == 0
	}
}

func main() {
	start := time.Now()

	mon_days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	cnt := 0

	year := 1901
	day := (1 + 365) % 7

	for year < 2001 {
		for i := 0; i < 12; i++ {
			if i == 1 {
				if isLeap(year) {
					day += 29
				} else {
					day += 28
				}
			} else {
				day += mon_days[i]
			}
			day %= 7
			if day == 0 {
				cnt++
			}
		}
		year++
	}

	fmt.Println(cnt)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
