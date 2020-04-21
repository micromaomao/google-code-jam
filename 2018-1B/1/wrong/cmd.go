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

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		test()
	}
}

type L struct {
	numVotes               int
	votesRequiredToRoundUp int
}

type Larr []L

func (l Larr) Len() int {
	return len(l)
}

func (l Larr) Less(i int, j int) bool {
	return l[i].votesRequiredToRoundUp < l[j].votesRequiredToRoundUp
}

func (l Larr) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func test() {
	var N, l int
	mustReadLineOfInts(&N, &l)
	Cis := mustReadLineOfIntsIntoArray()
	assert(len(Cis) == l)
	larr := make(Larr, l)
	votesLeft := N
	for i := 0; i < l; i++ {
		larr[i].numVotes = Cis[i]
		votesLeft -= larr[i].numVotes
	}
	assert(votesLeft > 0)
	for i := 0; i < l; i++ {
		var nvAdd = sort.Search(votesLeft+2, func(nvAdd int) bool {
			f := float64(larr[i].numVotes+nvAdd) / float64(N)
			return Round(f*100) > (f * 100)
		})
		larr[i].votesRequiredToRoundUp = nvAdd
	}
	sort.Sort(larr)
	for i := 0; i < l; i++ {
		if larr[i].votesRequiredToRoundUp <= votesLeft {
			larr[i].numVotes += larr[i].votesRequiredToRoundUp
			votesLeft -= larr[i].votesRequiredToRoundUp
			larr[i].votesRequiredToRoundUp = 0
		} else {
			break
		}
	}
	for votesLeft > 0 {
		useVotes := sort.Search(votesLeft+2, func(vUse int) bool {
			f := float64(vUse) / float64(N)
			return Round(f*100) > f*100
		})
		if useVotes > votesLeft {
			break
		}
		larr = append(larr, L{numVotes: useVotes})
		votesLeft -= useVotes
	}
	if votesLeft > 0 {
		larr = append(larr, L{numVotes: votesLeft})
	}
	sum := 0
	var realSum float64 = 0
	for i := 0; i < len(larr); i++ {
		f := float64(larr[i].numVotes) / float64(N)
		precentage := int(Round(f * 100))
		sum += precentage
		realSum += f
	}
	assert(Abs(realSum-1) < 1e-7)
	fmt.Fprintf(stdout, "%d\n", sum)
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

