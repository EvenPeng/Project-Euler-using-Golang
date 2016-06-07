/*
	a^2 + b^2 = c^2
	a + b + c = 1000

	a + b = 1000 - c
	a^2 + 2ab + b^2 = 1000^2 - 2000c + c^2

	2ab = 1000^2 - 2000c

	c = ab / 1000 - 500
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	done := false
	for a := 3; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				done = true
				break
			}
		}

		if done {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
