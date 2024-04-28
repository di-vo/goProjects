package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	boardSize = 5
)

var (
	nums   []int
	boards [][]int
)

func partOne(scanner bufio.Scanner) {
	i := 0
	tmp := make([]int, 0)

	for scanner.Scan() {
		if i == 0 {
			for _, e := range strings.Split(scanner.Text(), ",") {
				val, _ := strconv.Atoi(e)
				nums = append(nums, val)
			}
		}

		if i > 1 {
			if scanner.Text() != "" {
				trimLine := strings.ReplaceAll(scanner.Text(), "  ", " ")

				for _, e := range strings.Split(trimLine, " ") {
					val, _ := strconv.Atoi(e)
					tmp = append(tmp, val)
				}
			} else {
				boards = append(boards, tmp)
				tmp = make([]int, 0)
			}
		}

		i++
	}

	boards = append(boards, tmp)

	for _, n := range nums {
		for _, b := range boards {
			for i, e := range b {
				if e == n {
					b[i] = -1
				}
			}

			if hasWon(b) {
				tmp := 0

				for _, e := range b {
					if e != -1 {
						tmp += e
					}
				}

				fmt.Printf("result: %d\n", tmp * n)
				return
			}
		}
	}
}

func hasWon(board []int) bool {
	marks := make([]int, 0)

	for i, e := range board {
		if e == -1 {
			marks = append(marks, i)
		}
	}

	horHits := 0
	vertHits := 0

	for i := range marks {
		if i+1 < len(marks) {
			if marks[i+1] == marks[i] + 1 && (marks[i] + 1) % boardSize != 0 {
				horHits++
			} else {
                horHits = 0
            }

            if horHits == boardSize - 1 {
                return true
            }
		}
	}

	for i := range marks {
		if i+1 < len(marks) {
			if marks[i+1] == marks[i] + boardSize && (marks[i] + 1) % boardSize != 0 {
                vertHits++
			} else {
                vertHits = 0
            }

            if vertHits == boardSize - 1 {
                return true
            }
		}
	}

	// fmt.Printf("board: %d, marks: %d, hor: %d, vert: %d\n", board, marks, horHits, vertHits)

	return false
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

	partOne(*scanner)
}
