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
	var N int
	mustReadLineOfInts(&N)
	bff := mustReadLineOfIntsIntoArray()
	assert(len(bff) == N)
	for i := 0; i < N; i++ {
		bff[i]--
	}

	largestCycleSize := 0
	isTwoSizedCycles := make([]bool, N)
	tailLengths := make([]int, N)
	visited := make([]bool, N)
	for i := 0; i < N; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		// start dfs on each vertex
		cur := i
		iter := 0
		len := 0
		newVisited := make([]bool, N)
		for {
			if bff[bff[cur]] == cur {
				isTwoSizedCycles[cur] = true
				isTwoSizedCycles[bff[cur]] = true
				if tailLengths[cur] < len {
					tailLengths[cur] = len
				}
				break
			}
			len++
			iter++
			cur = bff[cur]
			if cur == i {
				break
			}
			newVisited[cur] = true
			if iter > N+1 {
				len = 0
				break
			}
		}
		if largestCycleSize < len {
			largestCycleSize = len
		}
		if len > 0 {
			for i := 0; i < N; i++ {
				if newVisited[i] {
					visited[i] = true
				}
			}
		}
	}

	twoSizedCycleCircleSize := 0
	for i := 0; i < N; i++ {
		if isTwoSizedCycles[i] {
			twoSizedCycleCircleSize += tailLengths[i] + 1
		}
	}

	if twoSizedCycleCircleSize > largestCycleSize {
		fmt.Fprintf(stdout, "%d\n", twoSizedCycleCircleSize)
	} else {
		fmt.Fprintf(stdout, "%d\n", largestCycleSize)
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
