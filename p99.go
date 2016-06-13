package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d99")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	ans := 0
	value := float64(0)

	for i := 1; i <= 1000; i++ {
		scanner.Scan()
		base, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		exp, _ := strconv.Atoi(scanner.Text())
		nextValue := math.Log10(float64(base)) * float64(exp)
		if nextValue > value {
			ans = i
			value = nextValue
		}
	}

	fmt.Println(ans, value)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
