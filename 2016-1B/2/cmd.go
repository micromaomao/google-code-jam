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
	line := strings.Split(mustReadLine(), " ")
	assert(len(line) == 2)
	C := []byte(line[0])
	J := []byte(line[1])
	assert(len(C) == len(J))
	for i := 0; i < len(C); i++ {
		c := C[i]
		j := J[i]
		if c == '?' {
			if j == '?' {
				if i == 0 {
					C[i] = '0'
					J[i] = '0'
				} else {
					cPrefix := string(C[0:i])
					jPrefix := string(J[0:i])
					if cPrefix < jPrefix {
						C[i] = '9'
						J[i] = '0'
					} else if cPrefix > jPrefix {
						C[i] = '0'
						J[i] = '9'
					} else {
						C[i] = '0'
						J[i] = '0'
					}
				}
			} else {
				C[i] = j
			}
		} else if j == '?' {
			J[i] = c
		}
	}
	stdout.Write(C)
	stdout.WriteByte(' ')
	stdout.Write(J)
	stdout.WriteByte('\n')
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
