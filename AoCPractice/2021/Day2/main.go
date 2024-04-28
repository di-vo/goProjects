package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	horPos, depth, aim int
)

func partOne(scanner bufio.Scanner) {
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		e, _ := strconv.Atoi(tmp[1])

		switch tmp[0] {
		case "forward":
			horPos += e
		case "down":
			depth += e
		case "up":
			depth -= e
		}
	}
}

func partTwo(scanner bufio.Scanner) {
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		e, _ := strconv.Atoi(tmp[1])

		switch tmp[0] {
		case "forward":
			horPos += e
            depth += aim * e
		case "down":
			aim += e
		case "up":
			aim -= e
		}
	}
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

	fmt.Printf("%d\n", horPos*depth)
}
