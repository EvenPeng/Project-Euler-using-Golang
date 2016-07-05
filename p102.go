package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Det(A, B [2]int) float64 {
	return float64(A[0]*B[1] - A[1]*B[0])
}

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d102")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var A, B, C [2]int
	ans := 0

	for i := 0; i < 1000; i++ {
		scanner.Scan()
		A[0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		A[1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		B[0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		B[1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		C[0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		C[1], _ = strconv.Atoi(scanner.Text())

		B[0] -= A[0]
		B[1] -= A[1]
		C[0] -= A[0]
		C[1] -= A[1]

		a := -1 * Det(A, C) / Det(B, C)
		b := Det(A, B) / Det(B, C)

		if a > 0 && b > 0 && a+b < 1 {
			ans++
		}
	}

	fmt.Println(ans)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
