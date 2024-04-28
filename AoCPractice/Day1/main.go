package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	nums = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func main() {
	// opening the file
	fi, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}

	// closing the file at the end of the program
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// reading the file
	scanner := bufio.NewScanner(fi)
	out := 0

	for scanner.Scan() {
		var wordBuf []byte
		var tmp []int

		for _, e := range scanner.Text() {
			if unicode.IsDigit(e) {
				t, _ := strconv.Atoi(string(e))
				tmp = append(tmp, t)
			}

			wordBuf = append(wordBuf, byte(e))

			for key, val := range nums {
				if strings.Contains(string(wordBuf), key) {
					tmp = append(tmp, val)
					clear(wordBuf)
					break
				}
			}
		}

		if len(tmp) > 0 {
			tmpR := strconv.Itoa(tmp[0]) + strconv.Itoa(tmp[len(tmp)-1])
			tmpS, _ := strconv.Atoi(tmpR)
			out += tmpS
			//fmt.Printf("%s %d, %d\n", scanner.Text(), tmp, tmpS)
		}
	}

	fmt.Println(out)
}
