package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	t := 286
	currT := t * (t + 1) / 2
	p := 166
	currP := p * (3*p - 1) / 2
	h := 144
	currH := h * (2*h - 1)
	max := 0
	if currT > currP {
		max = currT
	} else {
		max = currP
	}
	if currH > max {
		max = currH
	}

	for currT != currP || currP != currH {
		for currT < max {
			t++
			currT = t * (t + 1) / 2
		}
		if currT > max {
			max = currT
		}
		for currP < max {
			p++
			currP = p * (3*p - 1) / 2
		}
		if currP > max {
			max = currP
		}
		for currH < max {
			h++
			currH = h * (2*h - 1)
		}
		if currH > max {
			max = currH
		}
	}

	fmt.Println(max)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
