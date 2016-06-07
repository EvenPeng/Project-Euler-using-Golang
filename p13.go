package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d13")
	reader := bufio.NewReader(f)

	sum := int64(0)
	num := int64(0)

	for i := 0; i < 100; i++ {
		text, _ := reader.ReadString('\n')
		num, _ = strconv.ParseInt(text[:12], 10, 64)
		sum += num
	}

	fmt.Println("Ans:", strconv.FormatInt(sum, 10)[:10])

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
