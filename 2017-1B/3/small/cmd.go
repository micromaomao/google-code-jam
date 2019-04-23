// Only works for small!
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
	var N, Q int
	mustReadLineOfInts(&N, &Q)
	assert(Q == 1)
	maxDist := make([]int, 0, N)
	speed := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var E, S int
		mustReadLineOfInts(&E, &S)
		maxDist = append(maxDist, E)
		speed = append(speed, S)
	}
	lengths := make([]int, 0, N-1)
	for i := 0; i < N; i++ {
		destsDist := mustReadLineOfIntsIntoArray()
		assert(len(destsDist) == N)
		for j := 0; j < N; j++ {
			if j == i+1 {
				lengths = append(lengths, destsDist[j])
			} else {
				assert(destsDist[j] == -1)
			}
		}
	}
	assert(len(lengths) == N-1)
	var start, dest int
	mustReadLineOfInts(&start, &dest)
	assert(start == 1)
	assert(dest == N)
	minTimes := make([]float64, N)
	for i := 0; i < N; i++ {
		minTimes[i] = -1
	}
	minTimes[0] = 0
	for reach := 0; reach < N; reach++ {
		for withHorseFrom := 0; withHorseFrom <= reach; withHorseFrom++ {
			if reach == withHorseFrom {
			} else {
				requiredDist := 0
				for i := withHorseFrom; i < reach; i++ {
					requiredDist += lengths[i]
				}
				if maxDist[withHorseFrom] < requiredDist {
					continue
				}
				time := minTimes[withHorseFrom] + float64(requiredDist)/float64(speed[withHorseFrom])
				if time < minTimes[reach] || minTimes[reach] == -1 {
					minTimes[reach] = time
				}
			}
		}
	}
	fmt.Fprintf(stdout, "%.12f\n", minTimes[N-1])
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
