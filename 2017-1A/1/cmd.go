package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d:\n", i+1)
		test()
	}
}

func test() {
	var R, C int
	mustReadLineOfInts(&R, &C)
	cake := make([][]byte, R)
	for r := 0; r < R; r++ {
		readIn := mustReadLine()
		row := make([]byte, C)
		for c := 0; c < C; c++ {
			row[c] = readIn[c]
		}
		cake[r] = row
	}
	var firstNonEmptyRow []byte = nil
	rowIsEmpty := make([]bool, R)
	for r := 0; r < R; r++ {
		rowIsAllEmpty := true
		row := cake[r]
		var lastChar byte = '?'
		var firstChar byte = '?'
		for c := 0; c < C; c++ {
			if row[c] == '?' {
				row[c] = lastChar
				continue
			}
			rowIsAllEmpty = false
			if firstChar == '?' {
				firstChar = row[c]
			}
			lastChar = row[c]
		}
		rowIsEmpty[r] = rowIsAllEmpty
		if rowIsAllEmpty {
			continue
		}
		for c := 0; c < C && row[c] == '?'; c++ {
			row[c] = firstChar
		}
		if firstNonEmptyRow == nil {
			firstNonEmptyRow = row
		}
	}
	if firstNonEmptyRow == nil {
		panic("All empty.")
	}
	lastNonEmptyRow := firstNonEmptyRow
	for r := 0; r < R; r++ {
		if rowIsEmpty[r] {
			cake[r] = lastNonEmptyRow
		} else {
			lastNonEmptyRow = cake[r]
		}
		stdout.WriteString(string(cake[r]))
		stdout.WriteByte('\n')
	}
}

/*********Start boilerplate***********/

var stdin *bufio.Reader
var stdout *bufio.Writer

func main() {
	readFrom := os.Stdin
	if len(os.Args) == 2 {
		var err error
		readFrom, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	if len(os.Args) > 2 {
		panic("Too much arguments.")
	}
	stdin = bufio.NewReader(readFrom)
	stdout = bufio.NewWriter(os.Stdout)
	start()
	stdout.Flush()
}

func mustReadLine() string {
	str, err := stdin.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str = strings.TrimRight(str, "\n")
	return str
}

func mustAtoi(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return i
}

func debug(str string) {
	stdout.Flush()
	os.Stderr.WriteString("[debug] " + str + "\n")
}

func mustReadLineOfInts(a ...*int) {
	line := mustReadLine()
	strs := strings.Split(line, " ")
	if len(strs) != len(a) {
		panic("Expected " + strconv.Itoa(len(a)) + " numbers, got " + strconv.Itoa(len(strs)) + ".")
	}
	for i := 0; i < len(a); i++ {
		(*a[i]) = mustAtoi(strs[i])
	}
}
