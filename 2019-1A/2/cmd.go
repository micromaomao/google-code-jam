package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func start() {
	var t, maxNights, maxGophers int
	mustReadLineOfInts(&t, &maxNights, &maxGophers)
	for i := 0; i < t; i++ {
		test(maxNights, maxGophers)
	}
}

func test(maxNights, maxGophers int) {
	validGopherAmounts := make([]int, 0, maxGophers)
	for i := 1; i <= maxGophers; i++ {
		validGopherAmounts = append(validGopherAmounts, i)
	}
	queriesRem := maxNights
	mod := 18
	for {
		if queriesRem == 0 {
			os.Exit(0)
		}
		queriesRem--
		modStr := strconv.Itoa(mod)
		for i := 0; i < 18; i++ {
			stdout.WriteString(modStr)
			if i < 17 {
				stdout.WriteByte(' ')
			}
		}
		stdout.WriteByte('\n')
		stdout.Flush()
		windmails := mustReadLineOfIntsIntoArray()
		if len(windmails) == 1 {
			os.Exit(0)
		}
		assert(len(windmails) == 18)
		sum := 0
		for _, w := range windmails {
			sum += w
		}
		newValidGopherAmounts := make([]int, 0, len(validGopherAmounts))
		for _, gophers := range validGopherAmounts {
			if gophers%mod == sum%mod {
				newValidGopherAmounts = append(newValidGopherAmounts, gophers)
			}
		}
		validGopherAmounts = newValidGopherAmounts
		if len(validGopherAmounts) == 1 {
			stdout.WriteString(strconv.Itoa(validGopherAmounts[0]))
			stdout.WriteByte('\n')
			stdout.Flush()
			var verdict int
			mustReadLineOfInts(&verdict)
			if verdict == -1 {
				os.Exit(0)
			}
			return
		} else if len(validGopherAmounts) == 0 {
			os.Exit(0)
		}
		if mod == 2 {
			mod = 18
		} else {
			mod--
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
