package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func start() {
	assert(len(numbers) == 10)
	assert(len(order) == 10)
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		test()
	}
}

var numbers []string = []string{"ZERO", "ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
var order []int = []int{0, 2, 4, 6, 8, 3, 5, 7, 1, 9}

func test() {
	line := mustReadLine()
	freqMap := make(map[byte]int)
	for i := 0; i < len(line); i++ {
		freqMap[line[i]]++
	}
	result := make([]byte, 0)
	for tryNumIndex := 0; tryNumIndex < len(order); tryNumIndex++ {
		tmpMap := make(map[byte]int)
		for k, v := range freqMap {
			tmpMap[k] = v
		}
		ok := true
		tryNum := order[tryNumIndex]
		numString := numbers[tryNum]
		for i := 0; i < len(numString); i++ {
			l := numString[i]
			if tmpMap[l] > 0 {
				tmpMap[l]--
			} else {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		freqMap = tmpMap
		result = append(result, strconv.Itoa(tryNum)[0])
		tryNumIndex--
	}
	for _, v := range freqMap {
		if v != 0 {
			lefts := make([]byte, 0)
			for k, v := range freqMap {
				if v != 0 {
					lefts = append(lefts, k)
				}
			}
			panic(fmt.Errorf("Error: %v left, Current buffer is %v.", string(lefts), string(result)))
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	stdout.WriteString(string(result))
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
