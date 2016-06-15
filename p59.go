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

	f, _ := os.Open("./data/d59")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	alphabet := make([]rune, 26)
	for i := 0; i < 26; i++ {
		alphabet[i] = rune('a' + i)
	}

	mesg := []rune{}
	key := make([]rune, 3)
	percentage := float64(0)

	for scanner.Scan() {
		text := scanner.Text()
		ascii, _ := strconv.Atoi(text)
		mesg = append(mesg, rune(ascii))
	}

	for _, a := range alphabet {
		for _, b := range alphabet {
			for _, c := range alphabet {
				tmpKey := make([]rune, 3)
				tmpKey[0] = a
				tmpKey[1] = b
				tmpKey[2] = c

				cnt := float64(0)
				for i := 0; i < len(mesg); i++ {
					word := mesg[i] ^ tmpKey[i%3]
					if ('A' <= word && word <= 'Z') || ('a' <= word && word <= 'z') {
						cnt++
					}
				}

				if cnt/float64(len(mesg)) > percentage {
					key = tmpKey
					percentage = cnt / float64(len(mesg))
				}
			}
		}
	}

	for i := 0; i < len(mesg); i++ {
		mesg[i] ^= key[i%3]
	}

	fmt.Println("Key:", string(key))
	fmt.Println("Mesg:", string(mesg))

	ans := 0

	for i := 0; i < len(mesg); i++ {
		ans += int(mesg[i])
	}

	fmt.Println(ans)

	f.Close()

	elapsed := time.Since(start)
	fmt.Println("Elasped:", elapsed)
}
