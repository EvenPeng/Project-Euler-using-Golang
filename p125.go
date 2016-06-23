package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func IsPalindromes(x int) bool {
	str := strconv.FormatInt(int64(x), 10)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	limit := 100000000
	square := make([]int, limit/2+1)
	accu := make([]int, limit/2+1)
	elem := []int{}

	square[0] = 0
	accu[0] = 0

	for i := 1; i < limit/2+1; i++ {
		square[i] = i * i
		accu[i] = accu[i-1] + square[i]
	}

	for i := 0; i < limit/2+1; i++ {
		for j := i + 2; j < limit/2+1; j++ {
			curr := accu[j] - accu[i]
			if curr >= limit {
				break
			}
			if IsPalindromes(curr) {
				elem = append(elem, curr)
			}
		}
	}

	sort.Ints(elem)

	ans := elem[0]

	for i := 1; i < len(elem); i++ {
		if elem[i] != elem[i-1] {
			ans += elem[i]
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
