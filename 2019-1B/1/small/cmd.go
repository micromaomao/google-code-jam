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
	grid := make([][]int, size)
	for y := 0; y < size; y++ {
		grid[y] = make([]int, size)
	}
	for i := 0; i < numPeople; i++ {
		line := strings.Split(mustReadLine(), " ")
		assert(len(line) == 3)
		x := mustAtoi(line[0])
		y := mustAtoi(line[1])
		assert(len(line[2]) == 1)
		d := line[2][0]
		switch d {
		case 'N':
			for y := y + 1; y < size; y++ {
				for x := 0; x < size; x++ {
					grid[y][x]++
				}
			}
		case 'S':
			for y := y - 1; y >= 0; y-- {
				for x := 0; x < size; x++ {
					grid[y][x]++
				}
			}
		case 'E':
			for x := x + 1; x < size; x++ {
				for y := 0; y < size; y++ {
					grid[y][x]++
				}
			}
		case 'W':
			for x := x - 1; x >= 0; x-- {
				for y := 0; y < size; y++ {
					grid[y][x]++
				}
			}
		}
	}
	var maxX, maxY, maxN int
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if maxN < grid[y][x] {
				maxN = grid[y][x]
				maxX = x
				maxY = y
			}
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
