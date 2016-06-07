package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	maxD, maxLen := 7, 6

	for d := 11; d < 1000; d++ {
		q, r := []int{}, []int{}
		currQ, currR := 0, 0
		n := 1
		for n > 0 {
			n = (n % d) * 10
			currQ = n / d
			currR = n % d

			for i := len(q) - 1; i > -1; i-- {
				if q[i] == currQ && r[i] == currR {
					if maxLen < len(q)-i {
						maxLen = len(q) - i
						maxD = d
					}
					n = 0
					break
				}
			}

			q = append(q, currQ)
			r = append(r, currR)
		}
	}

	fmt.Println(maxD)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
