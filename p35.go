package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	isPrime := make([]bool, 1000000)
	for i := 0; i < 1000000; i += 2 {
		isPrime[i] = false
	}
	for i := 1; i < 1000000; i += 2 {
		isPrime[i] = true
	}
	isPrime[1] = false
	isPrime[2] = true

	for i := 3; i < 1000000; i += 2 {
		if isPrime[i] {
			for j := i * 2; j < 1000000; j += i {
				isPrime[j] = false
			}
		}
	}

	cnt := 1

	for curr := 3; curr < 1000000; curr += 2 {
		if isPrime[curr] {
			isCircle := true
			alias := curr
			aliasStr := []rune(strconv.FormatInt(int64(alias), 10))
			alias = 0
			for _, w := range aliasStr[1:] {
				alias = alias*10 + int(w-'0')
				if w == '0' {
					isCircle = false
				}
			}
			alias = alias*10 + int(aliasStr[0]-'0')

			for isCircle && alias != curr {
				if !isPrime[alias] {
					isCircle = false
				}
				aliasStr := []rune(strconv.FormatInt(int64(alias), 10))
				alias = 0
				for _, w := range aliasStr[1:] {
					alias = alias*10 + int(w-'0')
				}
				alias = alias*10 + int(aliasStr[0]-'0')
			}

			if isCircle {
				cnt++
			}
		}
	}

	fmt.Println(cnt)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
