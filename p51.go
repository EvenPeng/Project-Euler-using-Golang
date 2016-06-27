package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func IsPrime(prime []int64, x int64) bool {
	bound := int64(math.Sqrt(float64(x)))
	if x%bound == 0 {
		return false
	}

	for _, p := range prime {
		if p > bound {
			return true
		}
		if x%p == 0 {
			return false
		}
	}

	return true
}

func Add(set map[string][]rune, str []rune, d rune, beg int) {
	for i := beg; i < len(str); i++ {
		if str[i] == d {
			str[i] = rune('*')
			set[string(str)] = append(set[string(str)], d)
			Add(set, str, d, i+1)
			str[i] = d
		}
	}
}

func main() {
	start := time.Now()

	limit := int64(100)
	digits := []rune("0123456789")
	prime, maxP := []int64{2, 3, 5}, int64(7)
	done := false

	for maxP < int64(10) {
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
		}
		maxP += 2
	}

	for !done {
		set := make(map[string][]rune)

		for maxP < limit {
			if IsPrime(prime, maxP) {
				prime = append(prime, maxP)

				currStr := []rune(strconv.FormatInt(maxP, 10))
				for _, d := range digits {
					cnt := 0
					for _, td := range currStr {
						if td == d {
							cnt++
						}
					}

					if cnt > 0 {
						Add(set, currStr, d, 0)
					}
				}
			}
			maxP += 2
		}

		limit *= 10

		for currStr, d := range set {
			if len(d) == 8 {
				tmpStr := []rune(currStr)
				for i, c := range tmpStr {
					if c == rune('*') {
						tmpStr[i] = d[0]
					}
				}

				fmt.Println(string(tmpStr))
				done = true
				break
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
