package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d22")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	sort.Sort(sort.StringSlice(list))

	sum := 0

	for pos, word := range list {
		tmp_sum := len(word)
		for _, x := range word {
			tmp_sum += int(x - 'A')
		}
		sum += (pos + 1) * tmp_sum
	}

	fmt.Println(sum)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
