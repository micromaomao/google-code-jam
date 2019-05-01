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
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		test()
	}
}

func test() {
	var numPeople, size int
	mustReadLineOfInts(&numPeople, &size)
	size++
	xs := make([]int, size)
	ys := make([]int, size)
	for i := 0; i < numPeople; i++ {
		line := strings.Split(mustReadLine(), " ")
		assert(len(line) == 3)
		x := mustAtoi(line[0])
		y := mustAtoi(line[1])
		assert(len(line[2]) == 1)
		d := line[2][0]
		switch d {
		case 'N':
			for i := y + 1; i < size; i++ {
				ys[i]++
			}
		case 'S':
			for i := y - 1; i >= 0; i-- {
				ys[i]++
			}
		case 'E':
			for i := x + 1; i < size; i++ {
				xs[i]++
			}
		case 'W':
			for i := x - 1; i >= 0; i-- {
				xs[i]++
			}
		}
	}
	maxX := 0
	maxY := 0
	for i := 0; i < size; i++ {
		if xs[i] > xs[maxX] {
			maxX = i
		}
		if ys[i] > ys[maxY] {
			maxY = i
		}
	}
	fmt.Fprintf(stdout, "%d %d\n", maxX, maxY)
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
	defer stdout.Flush()
	start()
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

func mustReadLineOfIntsIntoArray() []int {
	line := mustReadLine()
	strs := strings.Split(line, " ")
	res := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = mustAtoi(strs[i])
	}
	return res
}

func assert(t bool) {
	if !t {
		panic("Assertion failed.")
	}
}
