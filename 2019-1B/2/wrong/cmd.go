/*
	For some reason, this code REs on Google's server and WAs on my computer with the testing tool.
*/

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
	var t, w int
	mustReadLineOfInts(&t, &w)
	for i := 0; i < t; i++ {
		test(w)
	}
}

var solutionMatrix [][]float64 = [][]float64{{1.0 / 10.0, -1.0 / 10.0, -1.0 / 20.0, 0, 0, 1.0 / 40.0}, {-6.0 / 5.0, 6.0 / 5.0, 1.0 / 10.0, 0, 0, -1.0 / 20.0}, {-2.0 / 5.0, -3.0 / 5.0, 6.0 / 5.0, 0, 0, -1.0 / 10.0}, {8.0 / 5.0, -8.0 / 5.0, -4.0 / 5.0, 1, 0, -1.0 / 10.0}, {-8.0 / 5.0, 8.0 / 5.0, 4.0 / 5.0, -1, 1, -2.0 / 5.0}, {12.0 / 5.0, -2.0 / 5.0, -6.0 / 5.0, 0, -1, 3.0 / 5.0}}

func Round(f float64) float64 {
	intPart := int64(f)
	f -= float64(intPart)
	if f >= 0.5 {
		return float64(intPart + 1)
	} else {
		return float64(intPart)
	}
}

func Abs(f float64) float64 {
	if f < 0 {
		return -f
	} else {
		return f
	}
}

func test(w int) {
	days := make([]int, 6)
	if w == 6 {
		for i := 1; i <= 6; i++ {
			fmt.Fprintf(stdout, "%d\n", i)
			stdout.Flush()
			var response int
			mustReadLineOfInts(&response)
			if response == -1 {
				os.Exit(1)
			}
			days[i-1] = response
		}
		solution := make([]int, 6)
		for i := 0; i < 6; i++ {
			row := solutionMatrix[i]
			var sum float64
			for j := 0; j < 6; j++ {
				sum += row[j] * float64(days[j])
			}
			solution[i] = int(Round(sum))
		}
		solutionStrings := make([]string, 6)
		for i := 0; i < 6; i++ {
			solutionStrings[i] = strconv.Itoa(solution[i])
		}
		stdout.WriteString(strings.Join(solutionStrings, " "))
		stdout.WriteByte('\n')
		stdout.Flush()
		var response int
		mustReadLineOfInts(&response)
		if response != 1 {
			os.Exit(1)
		}
	} else if w == 2 {
		stdout.WriteString("200\n")
		stdout.Flush()
		line := mustReadLine()
		rspTo200, _ := strconv.ParseUint(line, 10, 64)
		stdout.WriteString("55\n")
		stdout.Flush()
		line = mustReadLine()
		rspTo55, _ := strconv.ParseUint(line, 10, 64)
		a6 := (rspTo200 >> 33) % 128
		a5 := (rspTo200 >> 40) % 128
		a4 := (rspTo200 >> 50) % 128
		a1 := (rspTo55 >> 55) % 128
		a2 := (rspTo55 >> 27) % 128
		a3 := (rspTo55 >> 18) % 128
		fmt.Fprintf(stdout, "%d %d %d %d %d %d\n", int(a1), int(a2), int(a3), int(a4), int(a5), int(a6))
		stdout.Flush()
		var response int
		mustReadLineOfInts(&response)
		if response != 1 {
			os.Exit(1)
		}
	} else {
		panic(w)
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
