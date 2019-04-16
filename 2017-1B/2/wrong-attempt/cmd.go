// Fails with 4 2 0 2 0 0 0 -> RYRY
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
	var N, R, O, Y, G, B, V int
	mustReadLineOfInts(&N, &R, &O, &Y, &G, &B, &V)
	circle := make([]byte, 0, N)
	debug(string(circle))
	flag := false
	for {
		madeChanges := false
		if R > 0 {
			if len(circle) == 0 {
				circle = append(circle, 'R')
				R--
				madeChanges = true
			} else {
				for i := 0; i < len(circle); i++ {
					l := circle[i]
					r := circle[(i+1)%len(circle)]
					if l != 'R' && l != 'O' && l != 'V' && r != 'R' && r != 'O' && r != 'V' {
						if i == len(circle)-1 {
							circle = append(circle, 'R')
						} else {
							circle = append(circle, '_')
							copy(circle[i+2:], circle[i+1:len(circle)-1])
							circle[i+1] = 'R'
						}
						R--
						madeChanges = true
						break
					}
				}
			}
		}
		if O > 0 {
			if len(circle) == 0 {
				circle = append(circle, 'O')
				O--
				madeChanges = true
			}
		}
		if Y > 0 {
			if len(circle) == 0 {
				circle = append(circle, 'Y')
				Y--
				madeChanges = true
			} else {
				for i := 0; i < len(circle); i++ {
					l := circle[i]
					r := circle[(i+1)%len(circle)]
					if l != 'Y' && l != 'O' && l != 'G' && r != 'Y' && r != 'O' && r != 'G' {
						if i == len(circle)-1 {
							circle = append(circle, 'Y')
						} else {
							circle = append(circle, '_')
							copy(circle[i+2:], circle[i+1:len(circle)-1])
							circle[i+1] = 'Y'
						}
						Y--
						madeChanges = true
						break
					}
				}
			}
		}
		if G > 0 {
			if len(circle) == 0 {
				circle = append(circle, 'G')
				G--
				madeChanges = true
			}
		}
		if B > 0 {
			if len(circle) == 0 {
				circle = append(circle, 'B')
				B--
				madeChanges = true
			} else {
				for i := 0; i < len(circle); i++ {
					l := circle[i]
					r := circle[(i+1)%len(circle)]
					if l != 'B' && l != 'G' && l != 'V' && r != 'B' && r != 'G' && r != 'V' {
						if i == len(circle)-1 {
							circle = append(circle, 'B')
						} else {
							circle = append(circle, '_')
							copy(circle[i+2:], circle[i+1:len(circle)-1])
							circle[i+1] = 'B'
						}
						B--
						madeChanges = true
						break
					}
				}
			}
		}
		if V > 0 {
			if len(circle) == 0 {
				circle = append(circle, 'V')
				V--
				madeChanges = true
			}
		}
		if !madeChanges {
			if flag {
				break
			} else {
				flag = true
			}
		} else {
			flag = false
		}
		debug(string(circle))
	}
	if len(circle) == N {
		stdout.WriteString(string(circle))
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
