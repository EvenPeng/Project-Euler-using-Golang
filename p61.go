package main

import (
	"fmt"
	"time"
)

func Trail(sets [][]int, check []bool, head, tail, left int) int {
	for i, used := range check {
		if !used {
			for _, v := range sets[i] {
				if v/100 == tail {
					if left > 1 {
						check[i] = true
						curr := Trail(sets, check, head, v%100, left-1)
						check[i] = false
						if curr > 0 {
							return curr + v
						}
					} else {
						if v%100 == head {
							return v
						}
					}
				}
			}
		}
	}

	return 0
}

func main() {
	start := time.Now()

	check := make([]bool, 6)
	sets := make([][]int, 6)
	for i := 0; i < 6; i++ {
		sets[i] = []int{}
	}

	for n := 10; ; n++ {
		curr := n * (n + 1) / 2
		if curr > 999 {
			if curr > 9999 {
				break
			} else if curr%100 > 9 {
				sets[0] = append(sets[0], curr)
			}
		}
	}

	for n := 10; ; n++ {
		curr := n * n
		if curr > 999 {
			if curr > 9999 {
				break
			} else if curr%100 > 9 {
				sets[1] = append(sets[1], curr)
			}
		}
	}

	for n := 10; ; n++ {
		curr := n * (3*n - 1) / 2
		if curr > 999 {
			if curr > 9999 {
				break
			} else if curr%100 > 9 {
				sets[2] = append(sets[2], curr)
			}
		}
	}

	for n := 10; ; n++ {
		curr := n * (2*n - 1)
		if curr > 999 {
			if curr > 9999 {
				break
			} else if curr%100 > 9 {
				sets[3] = append(sets[3], curr)
			}
		}
	}

	for n := 10; ; n++ {
		curr := n * (5*n - 3) / 2
		if curr > 999 {
			if curr > 9999 {
				break
			} else if curr%100 > 9 {
				sets[4] = append(sets[4], curr)
			}
		}
	}

	for n := 10; ; n++ {
		curr := n * (3*n - 2)
		if curr > 999 {
			if curr > 9999 {
				break
			} else if curr%100 > 9 {
				sets[5] = append(sets[5], curr)
			}
		}
	}

	check[0] = true
	for _, v := range sets[0] {
		curr := Trail(sets, check, v/100, v%100, 5)
		if curr > 0 {
			fmt.Println(curr + v)
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
