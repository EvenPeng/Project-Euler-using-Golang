package main

import (
	"fmt"
	"strconv"
	"time"
)

func DigitsSum(x int64) int64 {
	sum := int64(0)

	for x > 0 {
		sum += x % 10
		x /= 10
	}

	return sum
}

func main() {
	start := time.Now()

	list := []int64{0, 0, 16}
	n := 0

	for i := int64(3); i < 16; i++ {
		list = append(list, i)
		for list[i] < 16 {
			list[i] *= i
		}
		if DigitsSum(list[i]) == i {
			n++
			fmt.Println(n, list[i])
		}
	}

	for n < 30 {
		list[2] *= 2
		if DigitsSum(list[2]) == 2 {
			n++
			fmt.Println(n, list[2])
		}

		for i := int64(len(list)); i < list[2]; i++ {
			if int64(len(strconv.FormatInt(list[2], 10))*9) < i {
				break
			}
			list = append(list, i)
		}

		for i := int64(3); i < int64(len(list)); i++ {
			if list[i] < list[2] {
				list[i] *= i
				if DigitsSum(list[i]) == i {
					n++
					fmt.Println(n, list[i])
				}
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
