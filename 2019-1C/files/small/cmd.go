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
	var A int
	mustReadLineOfInts(&A)
	otherProgs := make([][]byte, A)
	maxLen := A
	for i := 0; i < A; i++ {
		otherProgs[i] = []byte(mustReadLine())
	}
	myProg := make([]uint8, maxLen)
	options := []byte{'R', 'S', 'P'}
	for {
		won := true
	a:
		for _, opponent := range otherProgs {
			for i := 0; i < maxLen; i++ {
				oppAction := opponent[i%len(opponent)]
				myAction := options[myProg[i]]
				res := round(myAction, oppAction)
				if res == 1 {
					break
				}
				if res == -1 {
					won = false
					break a
				}
				if res == 0 && i == maxLen-1 {
					won = false
					break a
				}
			}
		}
		if won {
			for _, i := range myProg {
				stdout.WriteByte(options[i])
			}
			stdout.WriteByte('\n')
			return
		} else {
			// try next
			ptr := maxLen - 1
			for {
				if myProg[ptr]+1 >= 3 {
					myProg[ptr] = 0
					ptr--
					if ptr < 0 {
						stdout.WriteString("IMPOSSIBLE\n")
						return
					}
				} else {
					myProg[ptr]++
					break
				}
			}
		}
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
