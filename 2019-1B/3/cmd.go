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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func ConstructMaxTable(arr []int) [][]int {
	logArrLen := 0
	for arrLen := len(arr); arrLen > 0; arrLen >>= 1 {
		logArrLen++
	}
	maxTable := make([][]int, 0, logArrLen+1) // store indices, not values
	for currentSegLen := 1; currentSegLen <= len(arr); currentSegLen *= 2 {
		thisMaxRow := make([]int, len(arr)-currentSegLen+1)
		if currentSegLen == 1 {
			for i := 0; i < len(arr); i++ {
				thisMaxRow[i] = i
			}
		} else {
			lastMaxRow := maxTable[len(maxTable)-1]
			for i := 0; i < len(arr)-currentSegLen+1; i++ {
				indexA := lastMaxRow[i]
				indexB := lastMaxRow[i+currentSegLen/2]
				if arr[indexA] >= arr[indexB] { // perfer low index
					thisMaxRow[i] = indexA
				} else {
					thisMaxRow[i] = indexB
				}
			}
		}
		maxTable = append(maxTable, thisMaxRow)
	}
	return maxTable
}

// find the maximum element in range [start, end) and return its index, perfering lower index when equal.
func RangeMax(arr []int, maxTable [][]int, start, end int) int {
	rangeLen := uint(end - start)
	if rangeLen == 0 {
		return start
	}
	var rangeLevel uint = 0
	var rpower uint
	for rpower = 1; rpower < rangeLen; rpower <<= 1 {
		rangeLevel++
	}
	if (1 << rangeLevel) == rangeLen {
		return maxTable[rangeLevel][start]
	} else {
		rangeLevel--
		indexA := maxTable[rangeLevel][start]
		indexB := maxTable[rangeLevel][end-(1<<rangeLevel)]
		if arr[indexA] >= arr[indexB] {
			return indexA
		} else {
			return indexB
		}
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

var nTypes, K int

func test() {
	mustReadLineOfInts(&nTypes, &K)
	cSkills := mustReadLineOfIntsIntoArray()
	assert(len(cSkills) == nTypes)
	dSkills := mustReadLineOfIntsIntoArray()
	cMaxTable := ConstructMaxTable(cSkills)
	dMaxTable := ConstructMaxTable(dSkills)
	assert(len(dSkills) == nTypes)
	count := 0
	for i := 0; i < nTypes; i++ {
		currentChosenSkill := cSkills[i]
		// ! ! ! . . . ...
		stillChooseILeft := 0 // least index that still result in choosing i
		if i > 0 {
			stillChooseILeft = sort.Search(i, func(left int) bool {
				// still choose i?
				newMaxI := RangeMax(cSkills, cMaxTable, left, i)
				// Problem here:
				// return newMaxI <= currentChosenSkill
				return cSkills[newMaxI] <= currentChosenSkill
			})
		}
		goodEnoughLeft := sort.Search(i+1, func(left int) bool {
			dChoose := RangeMax(dSkills, dMaxTable, left, i+1)
			return dSkills[dChoose]-currentChosenSkill <= K
		})
		tooGoodLeft := sort.Search(i+1, func(left int) bool {
			dChoose := RangeMax(dSkills, dMaxTable, left, i+1)
			return !(currentChosenSkill-dSkills[dChoose] <= K)
		})
		if goodEnoughLeft == i+1 {
			continue
		}
		assert(goodEnoughLeft <= tooGoodLeft)
		stillChooseIRight := nTypes // first index that will result in not choosing i
		if i < nTypes-1 {
			stillChooseIRight = sort.Search(nTypes, func(right int) bool {
				if right <= i {
					return false
				}
				// Problem here:
				// newMaxI := RangeMax(cSkills, cMaxTable, i+1, right)
				newMaxI := RangeMax(cSkills, cMaxTable, i+1, right+1)
				return !(cSkills[newMaxI] < currentChosenSkill)
			})
		}
		firstNotGoodEnoughRight := sort.Search(nTypes, func(right int) bool {
			if right < i {
				return false
			}
			dChoose := RangeMax(dSkills, dMaxTable, i, right+1)
			return !(dSkills[dChoose]-currentChosenSkill <= K)
		})
		if firstNotGoodEnoughRight == i {
			continue
		}
		firstNotTooGoodRight := sort.Search(nTypes, func(right int) bool {
			if right < i {
				return false
			}
			dChoose := RangeMax(dSkills, dMaxTable, i, right+1)
			return currentChosenSkill-dSkills[dChoose] <= K
		})
		assert(firstNotTooGoodRight <= firstNotGoodEnoughRight)
		ans := (i - max(stillChooseILeft, goodEnoughLeft) + 1) * (min(firstNotGoodEnoughRight, stillChooseIRight) - i)
		if tooGoodLeft < firstNotTooGoodRight {
			ans -= (i - max(stillChooseILeft, tooGoodLeft) + 1) * (min(firstNotTooGoodRight, stillChooseIRight) - i)
		}
		assert(ans >= 0)
		count += ans
	}
	fmt.Fprintf(stdout, "%d\n", count)
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
