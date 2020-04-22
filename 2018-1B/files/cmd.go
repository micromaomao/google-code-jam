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

type Sign struct {
	d               int
	a, b            int
	m, n            int
	mLen, nLen      int
	setWithThisMLen int
	setWithThisNLen int
}

func test() {
	var numSigns int
	mustReadLineOfInts(&numSigns)
	signs := make([]Sign, numSigns)
	for i := 0; i < numSigns; i++ {
		var d, a, b int
		mustReadLineOfInts(&d, &a, &b)
		m := d + a
		n := d - b
		signs[i] = Sign{d, a, b, m, n, 0, 0, 0, 0}
		// debug(fmt.Sprintf("m=%v, n=%v", m, n))
	}
	maxValidSetLen := 0
	maxSetCount := 0
	// mLen pass
	for i := 0; i < len(signs); {
		currentM := signs[i].m
		mLen := 0
		for ; i+mLen < len(signs) && signs[i+mLen].m == currentM; mLen++ {
		}
		for j := i; j < i+mLen; j++ {
			signs[j].mLen = mLen - (j - i)
		}
		i += mLen
	}
	// nLen pass
	for i := 0; i < len(signs); {
		currentN := signs[i].n
		nLen := 0
		for ; i+nLen < len(signs) && signs[i+nLen].n == currentN; nLen++ {
		}
		for j := i; j < i+nLen; j++ {
			signs[j].nLen = nLen - (j - i)
		}
		i += nLen
	}

	for i := 0; i < len(signs); i++ {
		startSign := &signs[i]
		var setSize, m, n int
		var mSet, nSet bool
		continueWith := func(checkM bool) int {
			initCheckM := checkM
			sLen := 0
			for i+sLen < len(signs) {
				s := &signs[i+sLen]
				if checkM {
					if s.setWithThisMLen != 0 {
						sLen = s.setWithThisMLen
						return sLen
					}
					if mSet && s.m != m {
						break
					} else if !mSet {
						m = s.m
						mSet = true
					}
					sLen += s.mLen
				} else {
					if s.setWithThisNLen != 0 {
						sLen = s.setWithThisNLen
						return sLen
					}
					if nSet && s.n != n {
						break
					} else if !nSet {
						n = s.n
						nSet = true
					}
					sLen += s.nLen
				}
				checkM = !checkM
			}
			// set setWithThisM/NLen
			checkM = initCheckM
			newSLen := sLen
			nextFlip := -1
			if checkM {
				nextFlip = i + signs[i].mLen
			} else {
				nextFlip = i + signs[i].nLen
			}
			for j := i; j < i+sLen; j++ {
				s := &signs[j]
				if j == nextFlip {
					checkM = !checkM
					if checkM {
						nextFlip += s.mLen
					} else {
						nextFlip += s.nLen
					}
				}
				if nextFlip >= i+sLen {
					break // ignore last segment.
				}
				if checkM {
					s.setWithThisMLen = newSLen
				} else {
					s.setWithThisNLen = newSLen
				}
				newSLen--
			}
			return sLen
		}
		m = startSign.m
		mSet = true
		nSet = false
		a := continueWith(true)
		n = startSign.n
		nSet = true
		mSet = false
		b := continueWith(false)
		setSize = max(a, b)
		if maxValidSetLen < setSize {
			maxValidSetLen = setSize
			maxSetCount = 1
		} else if maxValidSetLen == setSize {
			maxSetCount++
		}
	}
	// debug(fmt.Sprintf("%v", signs))
	fmt.Fprintf(stdout, "%d %d\n", maxValidSetLen, maxSetCount)
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
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
