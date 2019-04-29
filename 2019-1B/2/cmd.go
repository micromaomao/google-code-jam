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
	if w == 6 {
		days := make([]int, 6)
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
			os.Exit(0)
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
		a6 := (rspTo200 >> 33) % 128 // My pervious attempts had 127 in here, which led to confusingly wrong answer.
		a5 := (rspTo200 >> 40) % 128 // However the code is still wrong even after fixing this.
		a4 := (rspTo200 >> 50) % 128
		a1 := (rspTo55 >> 55) % 128
		a2 := (rspTo55 >> 27) % 128
		// a3 := ((rspTo55 >> 18) - (a4 >> 5)) % 128
		//  Then I realized this, but it is still wrong, because addition not only affect the bits being added, but it may also
		//  carry on to one more bit to the left. Therefore, one should only start "chopping bits off" once all arithmetics are
		//  done. The correct code should be:
		a3 := ((rspTo55 - (a4 << 13) - (a5 << 11) - (a6 << 9)) >> 18) % 128
		fmt.Fprintf(stdout, "%d %d %d %d %d %d\n", a1, a2, a3, a4, a5, a6)
		stdout.Flush()
		var response int
		mustReadLineOfInts(&response)
		if response != 1 {
			os.Exit(0)
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
