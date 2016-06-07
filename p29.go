package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	cnt := 0
	done := make([]bool, 101)
	check := make([]bool, 1000)

	for a := 2; a < 101; a++ {
		done[a] = false
	}

	for a := 2; a < 101; a++ {
		if !done[a] {
			for b := 2; b < 1000; b++ {
				check[b] = false
			}

			for nextA, base := a, 1; nextA < 101; nextA, base = nextA*a, base+1 {
				done[nextA] = true
				for i := 2; i < 101; i++ {
					check[i*base] = true
				}
			}

			for b := 2; b < 1000; b++ {
				if check[b] {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
