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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func test() {
	var nTypes, K int
	mustReadLineOfInts(&nTypes, &K)
	cSkills := mustReadLineOfIntsIntoArray()
	assert(len(cSkills) == nTypes)
	dSkills := mustReadLineOfIntsIntoArray()
	assert(len(dSkills) == nTypes)
	count := 0
	for i := 0; i < nTypes; i++ {
		currentCMax := cSkills[i]
		currentDMax := dSkills[i]
		for j := i + 1; j <= nTypes; j++ {
			currentCMax = max(currentCMax, cSkills[j-1])
			currentDMax = max(currentDMax, dSkills[j-1])
			if abs(currentCMax-currentDMax) <= K {
				count++
			}
		}
	}
	fmt.Fprintf(stdout, "%d\n", count)
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
