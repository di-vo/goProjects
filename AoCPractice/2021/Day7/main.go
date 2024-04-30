package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	positions, diffs []int
)

func partOne(scanner bufio.Scanner) {
	for scanner.Scan() {
		// only one line
		tmp := strings.Split(scanner.Text(), ",")

		for _, e := range tmp {
			val, _ := strconv.Atoi(e)
			positions = append(positions, val)
		}
	}

    for _, p := range positions {
        tmp := 0

        for _, e := range positions {
            tmp += abs(p - e)
        }

        diffs = append(diffs, tmp)
    }

    result := slices.Min(diffs)
    fmt.Printf("result: %d\n", result)
}

func partTwo(scanner bufio.Scanner) {
	for scanner.Scan() {
		// only one line
		tmp := strings.Split(scanner.Text(), ",")

		for _, e := range tmp {
			val, _ := strconv.Atoi(e)
			positions = append(positions, val)
		}
	}

    for i := range slices.Max(positions) {
        tmp := 0

        for _, e := range positions {
            tmp += fuelVal(abs((i - e)))
        }

        diffs = append(diffs, tmp)
    }

    result := slices.Min(diffs)
    fmt.Printf("result: %d\n", result)
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}

func fuelVal(x int) int {
    result := 0

    for ;x > 0; x-- {
        result += x
    }

    return result
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
