package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	fieldSize = 990
)

var (
	lines []VectorLine
	field []Vector
)

type Vector struct {
	x, y, val int
}

type VectorLine struct {
	start, end Vector
}

func partOne(scanner bufio.Scanner) {
	// get lines
	for scanner.Scan() {
		text := strings.ReplaceAll(scanner.Text(), " ", "")
		linesText := strings.Split(text, "->")

		x1, _ := strconv.Atoi(strings.Split(linesText[0], ",")[0])
		y1, _ := strconv.Atoi(strings.Split(linesText[0], ",")[1])
		x2, _ := strconv.Atoi(strings.Split(linesText[1], ",")[0])
		y2, _ := strconv.Atoi(strings.Split(linesText[1], ",")[1])

		if x1 == x2 {
			if y1 < y2 {
				lines = append(lines, VectorLine{
					start: Vector{x: x1, y: y1, val: 1},
					end:   Vector{x: x2, y: y2, val: 1},
				})
			} else {
				lines = append(lines, VectorLine{
					start: Vector{x: x2, y: y2, val: 1},
					end:   Vector{x: x1, y: y1, val: 1},
				})
            }
		} else if y1 == y2 {
			if x1 < x2 {
				lines = append(lines, VectorLine{
					start: Vector{x: x1, y: y1, val: 1},
					end:   Vector{x: x2, y: y2, val: 1},
				})
			} else {
				lines = append(lines, VectorLine{
					start: Vector{x: x2, y: y2, val: 1},
					end:   Vector{x: x1, y: y1, val: 1},
				})
            }
        }
	}

	// get field
	for i := 0; i < fieldSize; i++ {
		for j := 0; j < fieldSize; j++ {
			field = append(field, Vector{
				x: i,
				y: j,
			})
		}
	}

	for _, e := range lines {
		if e.start.x == e.end.x {
			for i := e.start.y; i <= e.end.y; i++ {
				idx := slices.IndexFunc(field, func(v Vector) bool { return v.x == e.start.x && v.y == i })

				if idx != -1 {
					field[idx].val++
				}
			}
		} else if e.start.y == e.end.y {
			for i := e.start.x; i <= e.end.x; i++ {
				idx := slices.IndexFunc(field, func(v Vector) bool { return v.x == i && v.y == e.start.y })

				if idx != -1 {
					field[idx].val++
				}
			}
        }
	}

    // get result
    result := 0
    for _, e := range field {
        if e.val >= 2 {
            result++
        }
    }

    fmt.Printf("result: %d\n", result)
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
