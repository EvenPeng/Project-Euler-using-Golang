package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(prime []int, x int) bool {
	bound := int(math.Sqrt(float64(x)))
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

func main() {
	start := time.Now()

	prime, maxP := []int{2, 3, 5, 7, 11, 13}, 13
	accu := []int{0, 2, 5, 10, 17, 28, 41}
	ans, maxCnt := 0, 0

	for maxP < 1000000 {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
			accu = append(accu, maxP+accu[len(accu)-1])
		}
	}

	for p := len(prime) - 1; p > maxCnt; p-- {
		for i, j := 0, 0; i <= p && p-j > maxCnt; {
			if accu[i]-accu[j] > prime[p] {
				j++
			} else if accu[i]-accu[j] < prime[p] {
				i++
			} else {
				if i-j > maxCnt {
					ans = prime[p]
					maxCnt = i - j
				}
				break
			}
		}
	}

	fmt.Println(ans)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
