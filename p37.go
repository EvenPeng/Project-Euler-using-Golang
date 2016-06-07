package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func IsPrime(prime []int, x int) bool {
	bound := int(math.Sqrt(float64(x)))
	if x%bound == 0 {
		return false
	}

	for _, p := range prime {
		if p >= bound {
			return true
		}
		if x%p == 0 {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13

	sum := 0
	cnt := 0

	for cnt < 11 {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
			str := strconv.Itoa(maxP)

			isTuncatable := true
			for i := 1; i < len(str); i++ {
				curr, _ := strconv.Atoi(str[i:])
				if prime[sort.SearchInts(prime, curr)] != curr {
					isTuncatable = false
					break
				}

				curr, _ = strconv.Atoi(str[:len(str)-i])
				if prime[sort.SearchInts(prime, curr)] != curr {
					isTuncatable = false
					break
				}
			}

			if isTuncatable {
				fmt.Println("#", cnt, maxP)
				cnt++
				sum += maxP
			}
		}
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
