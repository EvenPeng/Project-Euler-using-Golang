package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	limit := int64(100)
	curr := int64(10)
	done := false

	for !done {
		set := make(map[string][]int64)
		for ; curr < limit; curr++ {
			cube := curr * curr * curr
			tmpStr := []rune(strconv.FormatInt(cube, 10))
			for i := 0; i < len(tmpStr); i++ {
				for j := 0; j < len(tmpStr); j++ {
					if tmpStr[i] > tmpStr[j] {
						tmpStr[i], tmpStr[j] = tmpStr[j], tmpStr[i]
					}
				}
			}
			set[string(tmpStr)] = append(set[string(tmpStr)], cube)
		}
		limit *= 10

		for _, val := range set {
			if len(val) == 5 {
				fmt.Println(val[0])
				done = true
				break
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
