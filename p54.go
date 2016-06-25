package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Rank(card []rune, val map[rune]int) (int, int) {
	straight := true
	flush := true

	for i := 2; i < 10; i += 2 {
		if val[card[i]]-1 != val[card[i-2]] {
			straight = false
			break
		}
	}

	for i := 3; i < 10; i += 2 {
		if card[i] != card[i-2] {
			flush = false
			break
		}
	}

	if straight && flush {
		// Royal Flush
		if card[0] == rune('T') {
			return 9, val[card[8]]
		}

		// Straight Flush
		return 8, val[card[8]]
	}

	cnt := make([]int, len(val))
	for i := 0; i < 10; i += 2 {
		cnt[val[card[i]]]++
	}

	four := 0
	three := 0
	two := 0

	fourVal := 0
	threeVal := 0
	twoVal := 0

	for i, c := range cnt {
		if c == 4 {
			four++
			fourVal = i
		} else if c == 3 {
			three++
			threeVal = i
		} else if c == 2 {
			two++
			twoVal = i
		}
	}

	// Four of a Kind
	if four > 0 {
		return 7, fourVal
	}

	// Full House
	if three > 0 && two > 0 {
		return 6, threeVal
	}

	// Flush
	if flush {
		return 5, val[card[8]]
	}

	// Straight
	if straight {
		return 4, val[card[8]]
	}

	// Three of a Kind
	if three > 0 {
		return 3, threeVal
	}

	// Two Pair
	if two > 1 {
		return 2, 0
	}

	// Pair
	if two > 0 {
		return 1, twoVal
	}

	return 0, val[card[8]]
}

func main() {
	start := time.Now()

	f, _ := os.Open("./data/d54")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	val := make(map[rune]int)
	val[rune('2')] = 0
	val[rune('3')] = 1
	val[rune('4')] = 2
	val[rune('5')] = 3
	val[rune('6')] = 4
	val[rune('7')] = 5
	val[rune('8')] = 6
	val[rune('9')] = 7
	val[rune('T')] = 8
	val[rune('J')] = 9
	val[rune('Q')] = 10
	val[rune('K')] = 11
	val[rune('A')] = 12

	A := make([]rune, 10)
	B := make([]rune, 10)
	ans := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 5; j++ {
			scanner.Scan()
			text := scanner.Text()
			A[j*2] = rune(text[0])
			A[j*2+1] = rune(text[1])
		}

		for j := 0; j < 5; j++ {
			scanner.Scan()
			text := scanner.Text()
			B[j*2] = rune(text[0])
			B[j*2+1] = rune(text[1])
		}

		for i := 0; i < 5; i++ {
			for j := i + 1; j < 5; j++ {
				if val[A[i*2]] > val[A[j*2]] {
					A[i*2], A[j*2] = A[j*2], A[i*2]
					A[i*2+1], A[j*2+1] = A[j*2+1], A[i*2+1]
				}

				if val[B[i*2]] > val[B[j*2]] {
					B[i*2], B[j*2] = B[j*2], B[i*2]
					B[i*2+1], B[j*2+1] = B[j*2+1], B[i*2+1]
				}
			}
		}

		rankA, valA := Rank(A, val)
		rankB, valB := Rank(B, val)

		if rankA > rankB {
			ans++
		} else if rankA == rankB {
			if valA > valB {
				ans++
			} else if valA == valB {
				for i := 8; i > -1; i -= 2 {
					if val[A[i]] > val[B[i]] {
						ans++
						fmt.Println(string(A[i]), string(B[i]))
						break
					} else if val[A[i]] < val[B[i]] {
						break
					}
				}
			}
		}
	}

	fmt.Println(ans)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
