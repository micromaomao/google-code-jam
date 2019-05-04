// Works for: Test set 1 and 2

// This approach use binary search, answering the question "can we make n lead"?
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

type Metal struct {
	ing1, ing2 int
	hasAmount  uint
}

var metals []Metal
var totalSupply uint

func test() {
	var M int
	mustReadLineOfInts(&M)
	metals = make([]Metal, M)
	for i := 0; i < M; i++ {
		var ing1, ing2 int
		mustReadLineOfInts(&ing1, &ing2)
		metals[i].ing1 = ing1 - 1
		metals[i].ing2 = ing2 - 1
	}
	line := strings.Split(mustReadLine(), " ")
	assert(len(line) == M)
	for i := 0; i < M; i++ {
		ps, err := strconv.ParseUint(line[i], 10, 32)
		if err != nil {
			panic(err)
		}
		metals[i].hasAmount = uint(ps)
		totalSupply += metals[i].hasAmount
	}

	var l, r uint
	r = totalSupply + 1
	for r-l > 1 {
		m := (l + r) / 2
		if check(m) {
			l = m
		} else {
			r = m
		}
	}
	fmt.Fprintf(stdout, "%d\n", l)
}

func check(m uint) bool {
	want := make(map[int]uint)
	want[0] = m
	totalWanted := m
	for {
		unsatisfied := -1
		for want, amount := range want {
			if metals[want].hasAmount < amount {
				unsatisfied = want
				break
			}
		}
		if unsatisfied == -1 {
			return true
		}
		additionalWantedAmount := want[unsatisfied] - metals[unsatisfied].hasAmount
		assert(additionalWantedAmount > 0)
		want[unsatisfied] = want[unsatisfied] - additionalWantedAmount
		want[metals[unsatisfied].ing1] += additionalWantedAmount
		want[metals[unsatisfied].ing2] += additionalWantedAmount
		totalWanted += additionalWantedAmount
		if totalWanted > totalSupply {
			return false
		}
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
