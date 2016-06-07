package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func WordValue(str string) int {
	sum := 0

	for i := 0; i < len(str); i++ {
		sum += int(str[i]-'A') + 1
	}

	return sum
}

func IsTriangle(x int) bool {
	n := int(math.Sqrt(float64(x * 2)))

	return (2 * x) == (n+1)*n
}

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d42")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	ans := 0

	for scanner.Scan() {
		if IsTriangle(WordValue(scanner.Text())) {
			ans++
		}
	}

	fmt.Println(ans)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
