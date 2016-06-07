package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	sum := int64(0)

	for curr := int64(1); curr < 1000000; curr++ {
		isPalindrome := true
		decStr := strconv.FormatInt(curr, 10)
		decLen := len(decStr)
		for i := 0; i < decLen/2; i++ {
			if decStr[i] != decStr[decLen-1-i] {
				isPalindrome = false
				break
			}
		}
		if !isPalindrome {
			continue
		}
		binStr := strconv.FormatInt(curr, 2)
		binLen := len(binStr)
		for i := 0; i < binLen/2; i++ {
			if binStr[i] != binStr[binLen-1-i] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			sum += curr
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
