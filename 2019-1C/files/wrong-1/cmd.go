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
	defeated := make([]bool, A)
	for i := 0; i < maxLen; i++ {
		myActionCanBe := make(map[byte]bool)
		myActionCanBe['R'] = true
		myActionCanBe['P'] = true
		myActionCanBe['S'] = true
		for opponent, p := range otherPrograms {
			if defeated[opponent] {
				continue
			}
			if i < len(p) {
				theirAction := p[i]
				switch theirAction {
				case 'R':
					delete(myActionCanBe, 'S')
				case 'P':
					delete(myActionCanBe, 'R')
				case 'S':
					delete(myActionCanBe, 'P')
				default:
					panic(theirAction)
				}
			}
		}
		if len(myActionCanBe) == 0 {
			stdout.WriteString("IMPOSSIBLE\n")
			return
		} else {
			myAct := byte('_')
			if len(myActionCanBe) == 1 {
				for act, _ := range myActionCanBe {
					myAct = act
					break
				}
				myProg[i] = myAct
				for opponent, theirProg := range otherPrograms {
					if defeated[opponent] {
						continue
					}
					if theirProg[i%len(theirProg)] != myAct {
						defeated[opponent] = true
					}
				}
			}
		}
	}
	defeatedAll := true
	for _, d := range defeated {
		if !d {
			defeatedAll = false
			break
		}
	}
	if !defeatedAll {
		stdout.WriteString("IMPOSSIBLE\n")
	} else {
		stdout.Write(myProg)
		stdout.WriteByte('\n')
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
