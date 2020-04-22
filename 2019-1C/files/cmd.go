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
	var t, F int
	mustReadLineOfInts(&t, &F)
	for i := 0; i < t; i++ {
		test(F)
	}
}

const A = byte('A')

type Pib struct {
	id  int
	arr []byte
}

var fact [6]int = [6]int{1, 1, 2, 6, 24, 120}

func test(F int) {
	query := func(pos int) byte {
		fmt.Fprintf(stdout, "%d\n", pos+1)
		stdout.Flush()
		ret := mustReadLine()[0]
		if ret == 'N' {
			os.Exit(0)
		}
		return ret - A
	}
	leftToQuery := make([]Pib, 119)
	for i := 0; i < len(leftToQuery); i++ {
		leftToQuery[i].id = i
		leftToQuery[i].arr = make([]byte, 5)
	}
	for prefixKnown := 0; prefixKnown < 3; prefixKnown++ {
		// debug(fmt.Sprintf("%v %v", prefixKnown, leftToQuery))
		counts := make([]int, 5)
		for i := 0; i < len(leftToQuery); i++ {
			lt := &leftToQuery[i]
			res := query(lt.id*5 + prefixKnown)
			counts[res]++
			lt.arr[prefixKnown] = res
		}
		letterIs := byte(0)
		expectedCount := 120
		for i := 0; i <= prefixKnown; i++ {
			expectedCount /= (5 - i)
		}
		for i, cnt := range counts {
			if cnt == expectedCount-1 {
				letterIs = byte(i)
				break
			} else if cnt != expectedCount {
				assert(cnt == 0)
			}
		}
		newLeftToQuery := make([]Pib, 0, len(leftToQuery))
		for _, lt := range leftToQuery {
			if lt.arr[prefixKnown] == letterIs {
				newLeftToQuery = append(newLeftToQuery, lt)
			}
		}
		assert(len(newLeftToQuery) == expectedCount-1)
		leftToQuery = newLeftToQuery
	}
	assert(len(leftToQuery) == 1)
	// find out last 2 letter with 1 query
	rsp := query(leftToQuery[0].id*5 + 3)
	leftToQuery[0].arr[4] = rsp // arr[3] is not rsp
	missing := leftToQuery[0].arr
	// last 1 letter
	letterMissing := make(map[byte]bool)
	for b := byte(0); b < 5; b++ {
		letterMissing[b] = true
	}
	for _, l := range missing[0:3] {
		delete(letterMissing, l)
	}
	delete(letterMissing, missing[4])
	var msLetter byte
	for l, _ := range letterMissing {
		msLetter = l
		break
	}
	missing[3] = msLetter
	for _, b := range missing {
		stdout.WriteByte(b + A)
	}
	stdout.WriteByte('\n')
	stdout.Flush()
	if mustReadLine() != "Y" {
		os.Exit(0)
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
