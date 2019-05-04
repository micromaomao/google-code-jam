/*
Referring to small-1...

The first idea of improvement here is that, if we already have, let's say, 100 x and 100 y, and lead requires x and
y, we can directly duduct 100 from each of the supply of x and y, and add 100 to leadCollected, rather than calling
takeMetal 100 times.

The second idea here is to "cache" the result of going through the deep calling stack of `takeMetal`. For example, if
we had already found out that in order to get 1 we need 2 and 3, and in order to get 2 we need 4 and 5, next time we
can just say in order to get 1 we need 3, 4, and 5, one of each.

These two optimizations enables us to reduce needing to repeatly call takeMetal as much as possible...

But really, a better way to do it is the method in the analysis - maintain a optimal "recipe". Therefore this file
implements that approach.

Disclaimer: the above ideas are produced only after reading the analysis, in an attempt to learn how I can go about
improving an algorithm during a real contest.
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
	line := strings.Split(mustReadLine(), " ")
	assert(len(line) == M)
	var totalAmountHas uint = 0
	for i := 0; i < M; i++ {
		ps, err := strconv.ParseUint(line[i], 10, 32)
		if err != nil {
			panic(err)
		}
		metals[i].hasAmount = uint(ps)
		totalAmountHas += metals[i].hasAmount
	}
	recipe := make(map[int]uint) // recipe to make metal 0
	recipe[0] = 1                // requires 1 metal 0 to produce 1 metal 0 (initially)
	var zeroMetalsProduced uint
a:
	for {
		limitingMetals := make([]int, 0)
		var produceAmount uint = (1 << 32) - 1
		var totalNumRequired uint = 0
		for metalRequired, amountRequired := range recipe {
			totalNumRequired += amountRequired
			hasAmount := metals[metalRequired].hasAmount
			if hasAmount < amountRequired {
				limitingMetals = append(limitingMetals, metalRequired)
				produceAmount = 0
			} else {
				canProduceAmount := hasAmount / amountRequired
				if canProduceAmount < produceAmount {
					produceAmount = canProduceAmount
				}
			}
		}
		if totalNumRequired > totalAmountHas {
			break
		}
		if produceAmount == 0 {
			for _, mtLacking := range limitingMetals {
				metalLacking := metals[mtLacking]
				hasAmount := metalLacking.hasAmount
				additionalRequiredAmount := recipe[mtLacking] - hasAmount
				if hasAmount == 0 {
					delete(recipe, mtLacking)
				} else {
					recipe[mtLacking] = hasAmount
				}
				if metalLacking.ing1 == mtLacking {
					break a
				} else {
					recipe[metalLacking.ing1] += additionalRequiredAmount
				}
				if metalLacking.ing2 == mtLacking {
					break a
				} else {
					recipe[metalLacking.ing2] += additionalRequiredAmount
				}
			}
		} else {
			for metalRequired, amountRequired := range recipe {
				deduction := amountRequired * produceAmount
				assert(metals[metalRequired].hasAmount >= deduction)
				metals[metalRequired].hasAmount -= deduction
				totalAmountHas -= deduction
				assert(totalAmountHas >= 0)
			}
			zeroMetalsProduced += produceAmount
		}
	}
	fmt.Fprintf(stdout, "%d\n", zeroMetalsProduced)
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
