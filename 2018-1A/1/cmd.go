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
		stdout.WriteString(fmt.Sprintf("Case #%d: ", i+1))
		possible := test()
		if possible {
			stdout.WriteString("POSSIBLE\n")
		} else {
			stdout.WriteString("IMPOSSIBLE\n")
		}
	}
}

func test() bool {
	var r, c, h, v int
	mustReadLineOfInts(&r, &c, &h, &v)
	grid := make([][]bool, r)
	sumgrid := make([][]int, r)
	for i := 0; i < r; i++ {
		grid[i] = make([]bool, c)
		sumgrid[i] = make([]int, c)
	}
	for l := 0; l < r; l++ {
		line := mustReadLine()
		if len(line) != c {
			panic("Invalid input.")
		}
		for x := 0; x < c; x++ {
			grid[l][x] = (line[x] == '@')
		}
	}
	// first row sum
	if grid[0][0] {
		sumgrid[0][0] = 1
	}
	for x := 1; x < c; x++ {
		sumgrid[0][x] = sumgrid[0][x-1]
		if grid[0][x] {
			sumgrid[0][x]++
		}
	}
	for y := 1; y < r; y++ {
		sumgrid[y][0] = sumgrid[y-1][0]
		if grid[y][0] {
			sumgrid[y][0]++
		}
	}
	for y := 1; y < r; y++ {
		for x := 1; x < c; x++ {
			sumgrid[y][x] = sumgrid[y-1][x] + sumgrid[y][x-1] - sumgrid[y-1][x-1]
			if grid[y][x] {
				sumgrid[y][x]++
			}
		}
	}
	if sumgrid[r-1][c-1] == 0 {
		return true
	}
	segV := sumgrid[r-1][c-1] / (v + 1)
	segVRem := sumgrid[r-1][c-1] % (v + 1)
	segH := sumgrid[r-1][c-1] / (h + 1)
	segHRem := sumgrid[r-1][c-1] % (h + 1)
	if segVRem > 0 || segHRem > 0 {
		return false
	}
	vSegs := make([]int, 1, v+1)
	hSegs := make([]int, 1, h+1)
	vPtr := 0
	for cV := 1; cV < v+1; cV++ {
		for ; sumgrid[r-1][vPtr] < segV*cV; vPtr++ {
		}
		if sumgrid[r-1][vPtr] != segV*cV {
			return false
		}
		vSegs = append(vSegs, vPtr+1)
	}
	hPtr := 0
	for cH := 1; cH < h+1; cH++ {
		for ; sumgrid[hPtr][c-1] < segH*cH; hPtr++ {
		}
		if sumgrid[hPtr][c-1] != segH*cH {
			return false
		}
		hSegs = append(hSegs, hPtr+1)
	}
	debug(fmt.Sprintf("vSegs = %v, hSegs = %v", vSegs, hSegs))
	expectRegionSum := -1
	for cV := 0; cV < len(vSegs); cV++ {
		for cH := 0; cH < len(hSegs); cH++ {
			startX := vSegs[cV]
			startY := hSegs[cH]
			var endX, endY int // non-inclusive
			if cV == len(vSegs)-1 {
				endX = c
			} else {
				endX = vSegs[cV+1]
			}
			if cH == len(hSegs)-1 {
				endY = r
			} else {
				endY = hSegs[cH+1]
			}
			regionSum := sumgrid[endY-1][endX-1]
			if startX > 0 {
				regionSum -= sumgrid[endY-1][startX-1]
			}
			if startY > 0 {
				regionSum -= sumgrid[startY-1][endX-1]
			}
			if startX > 0 && startY > 0 {
				regionSum += sumgrid[startY-1][startX-1]
			}
			if expectRegionSum == -1 {
				expectRegionSum = regionSum
			} else if expectRegionSum != regionSum {
				return false
			}
		}
	}
	return true
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
