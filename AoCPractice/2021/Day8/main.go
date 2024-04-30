package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partOne(scanner bufio.Scanner) {
    result := 0

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), "|")
        part := strings.Split(line[1], " ")

        for _,e := range part[1:] {
            if len(e) != 5 && len(e) != 6 {
                result++
                fmt.Printf("part: %s\n", e)
            }
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
