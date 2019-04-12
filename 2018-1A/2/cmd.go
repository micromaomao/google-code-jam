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
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		stdout.WriteString(fmt.Sprintf("Case #%d: ", i+1))
		test()
	}
}

type Cashier struct {
	M int
	S int
	P int
}

func test() {
	var R, B, C int
	mustReadLineOfInts(&R, &B, &C)
	cashiers := make([]Cashier, 0, C)
	minimumP := -1
	maxTime := 0
	for i := 0; i < C; i++ {
		var M, S, P int
		mustReadLineOfInts(&M, &S, &P)
		cashiers = append(cashiers, Cashier{M, S, P})
		if minimumP == -1 {
			minimumP = P
		} else if minimumP > P {
			minimumP = P
		}
		for _, cashier := range cashiers {
			thisMaxTime := cashier.P + cashier.S*cashier.M
			if maxTime < thisMaxTime {
				maxTime = thisMaxTime
			}
		}
	}
	tLeast := sort.Search(maxTime+1, func(t int) bool { // canFinishWithin t
		if t < minimumP {
			return false
		}
		cashiersCanDo := make([]int, 0, len(cashiers))
		for _, cashier := range cashiers {
			if t < cashier.P {
				cashiersCanDo = append(cashiersCanDo, 0)
				continue
			}
			canDo := (t - cashier.P) / cashier.S
			if canDo > cashier.M {
				canDo = cashier.M
			}
			cashiersCanDo = append(cashiersCanDo, canDo)
		}
		sort.Ints(cashiersCanDo)
		sumCanDos := 0
		for useNCashiers := 1; useNCashiers <= len(cashiers) && useNCashiers <= R; useNCashiers++ {
			sumCanDos += cashiersCanDo[len(cashiersCanDo)-useNCashiers]
			if sumCanDos >= B {
				return true
			}
		}
		return false
	})
	stdout.WriteString(fmt.Sprintf("%d\n", tLeast))
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
	start()
	stdout.Flush()
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
