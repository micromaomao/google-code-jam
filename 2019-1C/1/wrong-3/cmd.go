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

func round(my, other byte) int8 {
	switch my {
	case 'R':
		if other == 'P' {
			return -1
		} else if other == 'S' {
			return 1
		}
	case 'P':
		if other == 'S' {
			return -1
		} else if other == 'R' {
			return 1
		}
	case 'S':
		if other == 'R' {
			return -1
		} else if other == 'P' {
			return 1
		}
	default:
		panic("!")
	}
	return 0
}

func test() {
	var A int
	mustReadLineOfInts(&A)
	otherPrograms := make([][]byte, A)
	maxLen := 0
	for i := 0; i < A; i++ {
		otherPrograms[i] = []byte(mustReadLine())
		if len(otherPrograms[i]) > maxLen {
			maxLen = len(otherPrograms[i])
		}
	}
	myProg := make([]byte, maxLen)
	var next func(pos int, _defeatedOpponents []bool) bool
	next = func(pos int, _defeatedOpponents []bool) bool {
		opToTry := []byte{'R', 'P', 'S'}
		opResult := make([]int8, 3)
		opResult_defeatedComponent := make([][]bool, 3)
		for i := 0; i < 3; i++ {
			opResult_defeatedComponent[i] = make([]bool, len(_defeatedOpponents))
			copy(opResult_defeatedComponent[i], _defeatedOpponents)
		}
		for i, myOp := range opToTry {
			defeatedAll := true
			lose := false
			for opponentId, oppProgram := range otherPrograms {
				if opResult_defeatedComponent[i][opponentId] {
					continue
				}
				otherOp := oppProgram[pos%len(oppProgram)]
				vsRes := round(myOp, otherOp)
				if vsRes == 1 {
					opResult_defeatedComponent[i][opponentId] = true
				} else {
					defeatedAll = false
				}
				if vsRes == -1 {
					lose = true
					break
				}
			}
			if defeatedAll {
				opResult[i] = 1
			} else {
				if lose {
					opResult[i] = -1
				} else {
					opResult[i] = 0
				}
			}
		}
		for i, result := range opResult {
			myProg[pos] = opToTry[i]
			if result == 1 {
				maxLen = pos + 1
				return true
			}
		}
		for i, result := range opResult {
			myProg[pos] = opToTry[i]
			if result == 0 {
				if pos == len(myProg)-1 {
					continue
				}
				if next(pos+1, opResult_defeatedComponent[i]) {
					return true
				}
			}
		}
		return false
	}
	if next(0, make([]bool, A)) {
		stdout.Write(myProg[:maxLen])
		stdout.WriteByte('\n')
	} else {
		stdout.WriteString("IMPOSSIBLE\n")
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
