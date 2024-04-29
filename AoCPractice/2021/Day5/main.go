package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
    fieldSize = 10
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

		if x1 == x2 || y1 == y2 {
			lines = append(lines, VectorLine{
                start: Vector{x: x1, y: y1, val: 1},
                end:   Vector{x: x2, y: y2, val: 1},
			})
		}
	}

    sort.Slice(lines, func(i, j int) bool {
        if lines[i].start.x == lines[j].start.x {
            //TODO: implement sorting 
        }
    })

    // get field
    for i := 0; i < fieldSize; i++ {
        for j := 0; j < fieldSize; j++ {
            field = append(field, Vector{
                x: i,
                y: j,
            })
        }
    }

    fmt.Printf("lines: %d\n", lines)

    for _, e := range lines {
        if e.start.x == e.end.x {
            fmt.Printf("x1: %d, y1: %d, x2: %d, y2: %d\n", e.start.x, e.start.y, e.end.x, e.end.y)
            for i := e.start.y; i <= e.end.y; i++ {
                idx := slices.IndexFunc(field, func(v Vector) bool { return v.x == e.start.x && v.y == i }) 

                if idx != -1 {
                    field[idx].val++
                }
            }
        }
    }

    fmt.Printf("field after: %d\n", field)
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
