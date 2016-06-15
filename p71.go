package main

import (
	"fmt"
	"time"
)

func GCD(x, y int) int {
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}

	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	start := time.Now()

	d := 1000000
	if d%2 == 0 {
		d--
	}
	n := d/2 - 1

	for n*7 >= d*3 {
		n--
	}

	gcd := GCD(d, n)
	ansN := n / gcd
	ansD := d / gcd

	fmt.Println(ansN, "/", ansD)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
