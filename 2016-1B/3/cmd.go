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

type Topic struct {
	left, right string
}

func test() {
	var N int
	mustReadLineOfInts(&N)
	topics := make([]Topic, 0)
	for i := 0; i < N; i++ {
		line := strings.Split(mustReadLine(), " ")
		assert(len(line) == 2)
		topics = append(topics, Topic{line[0], line[1]})
	}
	count := 0
	leftMap := make(map[string]int)
	rightMap := make(map[string]int)
	for i, t := range topics {
		if _, exist := leftMap[t.left]; exist {
			continue
		}
		if _, seen := rightMap[t.right]; seen {
			continue
		}
		leftMap[t.left] = i
		rightMap[t.right] = i
		count++
	}
	for _, t := range topics {
		if _, exist := rightMap[t.right]; exist {
			continue
		}
		count++
	}
	for _, t := range topics {
		if _, exist := leftMap[t.left]; exist {
			continue
		}
		count++
	}
	fmt.Fprintf(stdout, "%d\n", N-count)
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
