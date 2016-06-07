package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	x := int64(1)

	for {
		if len(strconv.FormatInt(x, 10)) == len(strconv.FormatInt(x*6, 10)) {
			isSame := true

			orgStr := []rune(strconv.FormatInt(x, 10))
			for i := 0; i < len(orgStr); i++ {
				for j := i + 1; j < len(orgStr); j++ {
					if orgStr[i] > orgStr[j] {
						orgStr[i], orgStr[j] = orgStr[j], orgStr[i]
					}
				}
			}

			for m := 2; isSame && m < 7; m++ {
				tmpStr := []rune(strconv.FormatInt(x*int64(m), 10))
				for i := 0; i < len(tmpStr); i++ {
					for j := i + 1; j < len(tmpStr); j++ {
						if tmpStr[i] > tmpStr[j] {
							tmpStr[i], tmpStr[j] = tmpStr[j], tmpStr[i]
						}
					}
				}

				for i := 0; i < len(orgStr); i++ {
					if orgStr[i] != tmpStr[i] {
						isSame = false
					}
				}
			}

			if isSame {
				break
			}
		}
		x++
	}

	fmt.Println(x)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
