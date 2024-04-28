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

                fmt.Printf("result: %d, %d\n", tmp, n)
                return
            }
        }
    }
}

func hasWon(board []int) bool {
    fmt.Printf("%d\n", board)
	lb := 0
	ub := lb + boardSize
    tmp := make([]int, 0)

	// horizontal check
	for ub+boardSize <= len(board) {
        fmt.Printf("%d\n", board[lb:ub])
		for _, e := range board[lb:ub] {
			if e == -1 {
                tmp = append(tmp, e)
			}
		}

		lb += boardSize
		ub += boardSize

        if len(tmp) == 5 {
            return true
        } else {
            tmp = make([]int, 0)
        }
	}

    tmp = make([]int, 0)

	// vertical check
	for range boardSize {
		var tmp []int

		val := 0
		for range boardSize {
			tmp = append(tmp, board[val])
			val += boardSize
		}

		for _, e := range tmp {
			if e == -1 {
			    tmp = append(tmp, e)
			}
		}

        if len(tmp) == 5 {
            return true
        } else {
            tmp = make([]int, 0)
        }
	}

    
    return false
}

func main() {
	fi, err := os.Open("input.txt")
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
