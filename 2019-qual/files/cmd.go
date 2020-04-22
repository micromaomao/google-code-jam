package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		var _N, _B, _F int
		mustReadLineOfInts(&_N, &_B, &_F)
		F := uint(_F)
		N := uint(_N)
		B := uint(_B)
		workerNumbers := make([]uint, N)
		for i := uint(0); i < N; i++ {
			workerNumbers[i] = i % (1 << F)
		}
		gotNumbers := make([]uint, N-B)
		for i := uint(0); i < F; i++ {
			for w := 0; w < len(workerNumbers); w++ {
				if (workerNumbers[w]>>i)&1 > 0 {
					stdout.WriteByte('1')
				} else {
					stdout.WriteByte('0')
				}
			}
			stdout.WriteByte('\n')
			stdout.Flush()
			response := mustReadLine()
			if response == "-1" {
				os.Exit(0)
			} else {
				assert(uint(len(response)) == N-B)
				for r := uint(0); r < N-B; r++ {
					switch response[r] {
					case '0':
						// do nothing
					case '1':
						gotNumbers[r] += 1 << i
					}
				}
			}
		}
		indices := make([]uint, 0, B)
		pointToRsp := 0
		for i := uint(0); i < N; i++ {
			if pointToRsp < len(gotNumbers) && gotNumbers[pointToRsp] == workerNumbers[i] {
				pointToRsp++
				continue
			} else {
				indices = append(indices, i)
			}
		}
		assert(len(indices) == int(B))
		for i, id := range indices {
			if i != 0 {
				stdout.WriteByte(' ')
			}
			stdout.WriteString(strconv.Itoa(int(id)))
		}
		stdout.WriteByte('\n')
		stdout.Flush()
		verdict := mustReadLine()
		if verdict == "-1" {
			os.Exit(0)
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
