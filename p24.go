package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	factorial := make([]int64, 10)
	factorial[1] = 1
	for i := 2; i < 10; i++ {
		factorial[i] = factorial[i-1] * int64(i)
	}

	seq := make([]int, 10)
	for i := 0; i < 10; i++ {
		seq[i] = i
	}

	diff := int64(999999)
	for diff > 0 {
		seqLen := len(seq)
		out := diff / factorial[seqLen-1]
		fmt.Print(seq[out])
		seq = append(seq[:int(out)], seq[int(out)+1:]...)
		diff -= out * factorial[seqLen-1]
	}

	for _, d := range seq {
		fmt.Print(d)
	}

	fmt.Println()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
