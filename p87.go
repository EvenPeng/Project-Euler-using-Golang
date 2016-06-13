package main

import (
	"fmt"
	"math"
	"sort"
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

	prime, maxP := []int{2, 3}, int(3)
	square := []int{4, 9}
	cubic := []int{8, 27}
	fourth := []int{16, 81}

	for square[len(square)-1]+cubic[0]+fourth[0] < 50000000 {
		maxP += 2
		if IsPrime(prime, maxP) {
			prime = append(prime, maxP)
			square = append(square, maxP*maxP)
		}
	}

	for square[0]+cubic[len(cubic)-1]+fourth[0] < 50000000 {
		cubic = append(cubic, square[len(cubic)]*prime[len(cubic)])
	}

	for square[0]+cubic[0]+fourth[len(fourth)-1] < 50000000 {
		fourth = append(fourth, cubic[len(fourth)]*prime[len(fourth)])
	}

	resultArr := []int{}

	for _, s := range square {
		for _, c := range cubic {
			for _, f := range fourth {
				if s+c+f < 50000000 {
					resultArr = append(resultArr, s+c+f)
				}
			}
		}
	}

	sort.IntSlice(resultArr).Sort()

	cnt := int(1)

	for i := 1; i < len(resultArr); i++ {
		if resultArr[i] != resultArr[i-1] {
			cnt++
		}
	}

	fmt.Println(cnt)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
