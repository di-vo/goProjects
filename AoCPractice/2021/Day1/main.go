package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(scanner bufio.Scanner) {
	lastVal, incs := 0, 0

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err == nil && lastVal > 0 && val > lastVal {
			incs++
		}

		lastVal = val
	}

	fmt.Println(incs)
}

func partTwo(scanner bufio.Scanner) {
	var vals []int
	incs := 0

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())

		if err == nil {
			vals = append(vals, val)
		}
	}

	lb := 0
	ub := lb + 3

	for {
		tmp1, tmp2 := 0, 0

		if ub+1 <= len(vals) {
			for _, e := range vals[lb:ub] {
				tmp1 += e
			}

			for _, e := range vals[lb+1 : ub+1] {
				tmp2 += e
			}
		} else {
			break
		}

        // fmt.Printf("%d, %d, %d\n", tmp1, tmp2, incs)
		if tmp1 < tmp2 {
			incs++
		}

        lb++
        ub++
	}

	fmt.Println(incs)
}

func main() {
	fi, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	//partOne(*scanner)
	partTwo(*scanner)
}
