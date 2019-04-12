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
		line := strings.Split(mustReadLine(), " ")
		sMax := mustAtoi(line[0])
		audArray := make([]int, 0, len(line[1]))
		for _, n := range line[1] {
			audArray = append(audArray, mustAtoi(string(n)))
		}
		var numInsertion int
		var numStoodUp int
		for sLevel, numPpl := range audArray {
			if sLevel <= numStoodUp {
				numStoodUp += numPpl
				if numStoodUp >= sMax {
					break
				}
			} else {
				numInsertion += sLevel - numStoodUp
				numStoodUp += sLevel - numStoodUp + numPpl
			}
			// debug(fmt.Sprintf("at sLevel = %v: numStoodUp = %v, numInsertion = %v", sLevel, numStoodUp, numInsertion))
		}
		stdout.WriteString(fmt.Sprintf("Case #%d: %d\n", i+1, numInsertion))
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
