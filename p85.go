package main

import (
	"fmt"
	"time"
)

func absInt(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}

func main() {
	start := time.Now()

	ansW, ansH, minDiff := 2000, 2000, 1999999

	for w := 1; w < ansH; w++ {
		rowCnt := (w + 1) * w / 2
		for h := ansH; h > 0; h-- {
			colCnt := (h + 1) * h / 2
			if absInt(rowCnt*colCnt-2000000) < minDiff {
				ansW = w
				ansH = h
				minDiff = absInt(rowCnt*colCnt - 2000000)
				fmt.Println(w, h, minDiff)
			}
		}
	}

	fmt.Println(ansW * ansH)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
