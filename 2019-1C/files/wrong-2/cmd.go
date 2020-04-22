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
	otherPrograms := make([][]byte, A)
	maxLen := 0
	for i := 0; i < A; i++ {
		otherPrograms[i] = []byte(mustReadLine())
		if len(otherPrograms[i]) > maxLen {
			maxLen = len(otherPrograms[i])
		}
	}
	myProg := make([]byte, maxLen)
	check := func(upTo int) int8 {
		defeated := make([]bool, A)
		for ps := 0; ps <= upTo; ps++ {
			allDefeated := true
			for op := 0; op < A; op++ {
				if defeated[op] {
					continue
				}
				opProg := otherPrograms[op]
				opAct := opProg[ps%len(opProg)]
				myAct := myProg[ps]
				switch myAct {
				case 'R':
					if opAct == 'P' {
						return -1
					} else if opAct == 'S' {
						defeated[op] = true
					}
				case 'P':
					if opAct == 'S' {
						return -1
					} else if opAct == 'R' {
						defeated[op] = true
					}
				case 'S':
					if opAct == 'R' {
						return -1
					} else if opAct == 'P' {
						defeated[op] = true
					}
				}
				if !defeated[op] {
					allDefeated = false
				}
			}
			if allDefeated {
				return 1
			}
		}
		return 0
	}
	var search func(currentUpTo int) bool
	search = func(currentUpTo int) bool {
		if currentUpTo == maxLen {
			return false
		}
		toTry := []byte{'R', 'P', 'S'}
		for _, try := range toTry {
			myProg[currentUpTo] = try
			ck := check(currentUpTo)
			if ck == 1 {
				maxLen = currentUpTo + 1
				return true
			} else if ck == -1 {
				continue
			} else {
				if search(currentUpTo + 1) {
					return true
				} // else continue
			}
		}
		return false
	}
	if search(0) {
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
