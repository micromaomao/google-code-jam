package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
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

type CJ struct {
	C, J []byte
}

func rank(a, b *CJ) *CJ {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	aC, ok := big.NewInt(0).SetString(string(a.C), 10)
	assert(ok)
	aJ, ok := big.NewInt(0).SetString(string(a.J), 10)
	assert(ok)
	bC, ok := big.NewInt(0).SetString(string(b.C), 10)
	assert(ok)
	bJ, ok := big.NewInt(0).SetString(string(b.J), 10)
	assert(ok)
	diffA := big.NewInt(0).Sub(aC, aJ)
	diffB := big.NewInt(0).Sub(bC, bJ)
	cmpAbs := diffA.CmpAbs(diffB)
	if cmpAbs < 0 {
		return a
	} else if cmpAbs > 0 {
		return b
	}
	cmpC := aC.Cmp(bC)
	if cmpC < 0 {
		return a
	} else if cmpC > 0 {
		return b
	}
	if aJ.Cmp(bJ) < 0 {
		return a
	} else {
		return b
	}
}

func assumeAndGoOn(pos int, C, J []byte, cmp int8) *CJ {
	nC := make([]byte, len(C))
	nJ := make([]byte, len(J))
	copy(nC, C)
	copy(nJ, J)
	C = nC
	J = nJ
	N := len(C)
	if cmp == 0 {
		for ; pos < N; pos++ {
			if C[pos] == '?' && J[pos] == '?' {
				C[pos] = '0'
				J[pos] = '0'
			} else if C[pos] == '?' && J[pos] != '?' {
				C[pos] = J[pos]
			} else if C[pos] != '?' && J[pos] == '?' {
				J[pos] = C[pos]
			}
		}
		return &CJ{C, J}
	} else if cmp < 0 {
		for ; pos < N; pos++ {
			if C[pos] == '?' {
				C[pos] = '9'
			}
			if J[pos] == '?' {
				J[pos] = '0'
			}
		}
		return &CJ{C, J}
	} else {
		for ; pos < N; pos++ {
			if C[pos] == '?' {
				C[pos] = '0'
			}
			if J[pos] == '?' {
				J[pos] = '9'
			}
		}
		return &CJ{C, J}
	}
}

func test() {
	line := strings.Split(mustReadLine(), " ")
	assert(len(line) == 2)
	C := []byte(line[0])
	J := []byte(line[1])
	assert(len(C) == len(J))
	N := len(C)
	var best *CJ = nil
	notEqual := false
	for i := 0; i < N; i++ {
		if C[i] == J[i] && C[i] != '?' {
			continue
		}
		if C[i] != '?' && J[i] != '?' {
			notEqual = true
			cmp := int8(0)
			if C[i] < J[i] {
				cmp = -1
			} else if C[i] > J[i] {
				cmp = 1
			}
			best = rank(best, assumeAndGoOn(i+1, C, J, cmp))
			break
		} else if C[i] != '?' && J[i] == '?' {
			if C[i] > '0' {
				J[i] = C[i] - 1
				best = rank(best, assumeAndGoOn(i+1, C, J, 1))
			}
			if C[i] < '9' {
				J[i] = C[i] + 1
				best = rank(best, assumeAndGoOn(i+1, C, J, -1))
			}
			J[i] = C[i]
			continue
		} else if C[i] == '?' && J[i] != '?' {
			if J[i] > '0' {
				C[i] = J[i] - 1
				best = rank(best, assumeAndGoOn(i+1, C, J, -1))
			}
			if J[i] < '9' {
				C[i] = J[i] + 1
				best = rank(best, assumeAndGoOn(i+1, C, J, 1))
			}
			C[i] = J[i]
			continue
		} else if C[i] == '?' && J[i] == '?' {
			C[i] = '1'
			J[i] = '0'
			best = rank(best, assumeAndGoOn(i+1, C, J, 1))
			C[i] = '0'
			J[i] = '1'
			best = rank(best, assumeAndGoOn(i+1, C, J, -1))
			C[i] = '0'
			J[i] = '0'
			continue
		}
		panic("Uncovered case!")
	}
	if !notEqual {
		best = rank(best, &CJ{C, J})
	}
	stdout.Write(best.C)
	stdout.WriteByte(' ')
	stdout.Write(best.J)
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
