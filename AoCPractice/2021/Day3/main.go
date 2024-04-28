package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	arrSize = 12
)

var (
	zeros, ones, gammaBin, epsBin [arrSize]int
	oxElems, scrubElems           []string
)

func partOne(scanner bufio.Scanner) {
	for scanner.Scan() {
		for i, e := range scanner.Text() {
			if n, _ := strconv.Atoi(string(e)); n == 0 {
				zeros[i]++
			} else {
				ones[i]++
			}
		}
	}

	for i := range zeros {
		if zeros[i] > ones[i] {
			epsBin[i] = 1
		} else {
			gammaBin[i] = 1
		}
	}
	gamStr, epsStr := "", ""

	for _, e := range gammaBin {
		tmp := strconv.Itoa(e)
		gamStr += tmp
	}

	for _, e := range epsBin {
		tmp := strconv.Itoa(e)
		epsStr += tmp
	}

	gamDec, _ := strconv.ParseInt(gamStr, 2, 64)
	epsDec, _ := strconv.ParseInt(epsStr, 2, 64)

	fmt.Printf("result: %d\n", gamDec*epsDec)
}

func partTwo(scanner bufio.Scanner) {
	for scanner.Scan() {
		oxElems = append(oxElems, scanner.Text())
		scrubElems = append(scrubElems, scanner.Text())
	}

	i := 0
	currentMax := 0

	for len(oxElems) > 1 {
		for id := range zeros {
			zeros[id] = 0
			ones[id] = 0
		}

		for _, e := range oxElems {
			for id, c := range e {
				if n, _ := strconv.Atoi(string(c)); n == 0 {
					zeros[id]++
				} else {
					ones[id]++
				}
			}
		}

		var tmpElems []string

		if zeros[i] > ones[i] {
			currentMax = 0
		} else {
			currentMax = 1
		}

		for _, e := range oxElems {
			if val, _ := strconv.Atoi(string(e[i])); val == currentMax {
				tmpElems = append(tmpElems, e)
			}
		}

		// fmt.Printf("%s, max: %d, i: %d, zeros: %d, ones: %d\n", oxElems, currentMax, i, zeros[i], ones[i])
		oxElems = tmpElems

		if i == arrSize-1 {
			i = 0
		} else {
			i++
		}
	}

    i = 0
    currentMin := 0

	for len(scrubElems) > 1 {
		for id := range zeros {
			zeros[id] = 0
			ones[id] = 0
		}

		for _, e := range scrubElems {
			for id, c := range e {
				if n, _ := strconv.Atoi(string(c)); n == 0 {
					zeros[id]++
				} else {
					ones[id]++
				}
			}
		}

		var tmpElems []string

		if zeros[i] > ones[i] {
			currentMin = 1
		} else {
			currentMin = 0
		}

		for _, e := range scrubElems {
			if val, _ := strconv.Atoi(string(e[i])); val == currentMin {
				tmpElems = append(tmpElems, e)
			}
		}

		// fmt.Printf("%s, max: %d, i: %d, zeros: %d, ones: %d\n", scrubElems, currentMax, i, zeros[i], ones[i])
		scrubElems = tmpElems

		if i == arrSize-1 {
			i = 0
		} else {
			i++
		}
	}

    oxStr := oxElems[0]
    scrubStr := scrubElems[0]

	oxDec, _ := strconv.ParseInt(oxStr, 2, 64)
	scrubDec, _ := strconv.ParseInt(scrubStr, 2, 64)

    fmt.Printf("result: %d\n", oxDec * scrubDec)
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

	// partOne(*scanner)
	partTwo(*scanner)
}
