package main

import (
	"fmt"
	"sort"
	"time"
)

type Rad struct {
	v, e int
}

type RadSlice []Rad

func (r RadSlice) Len() int {
	return len(r)
}

func (r RadSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RadSlice) Less(i, j int) bool {
	if r[i].e == r[j].e {
		return r[i].v < r[j].v
	} else {
		return r[i].e < r[j].e
	}
}

func main() {
	start := time.Now()

	limit := 100001
	E := make([]Rad, limit)

	for i := 0; i < limit; i++ {
		E[i].v = i
		E[i].e = 1
	}

	for i := 2; i < limit; i++ {
		if E[i].e == 1 {
			for j := i; j < limit; j += i {
				E[j].e *= i
			}
		}
	}

	sort.Sort(RadSlice(E))

	fmt.Println(E[10000].v)

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
