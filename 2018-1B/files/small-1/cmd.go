// Works for: Test set 1

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
	hasAmounts := mustReadLineOfIntsIntoArray()
	assert(len(hasAmounts) == M)
	for i := 0; i < M; i++ {
		metals[i].hasAmount = uint(hasAmounts[i])
	}
	collectedZeroMetals := 0
	for {
		has := takeMetal(0, 0)
		if has {
			collectedZeroMetals++
		} else {
			break
		}
	}
	fmt.Fprintf(stdout, "%d\n", collectedZeroMetals)
}

type ExclusionList map[int]bool

func takeMetal(metal int, depth int) bool {
	target := &metals[metal]
	if target.hasAmount > 0 {
		target.hasAmount--
		return true
	}
	if depth > len(metals) {
		return false
	}
	tookIg1 := takeMetal(target.ing1, depth+1)
	tookIg2 := takeMetal(target.ing2, depth+1)
	if tookIg1 && tookIg2 {
		return true
	}
	if tookIg1 {
		metals[target.ing1].hasAmount++
	}
	if tookIg2 {
		metals[target.ing2].hasAmount++
	}
	return false
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
