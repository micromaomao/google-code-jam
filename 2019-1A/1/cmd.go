package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		var R, C int
		mustReadLineOfInts(&R, &C)
		test(R, C)
	}
}

type Point struct {
	r, c int
}

func printGrid(grid [][]bool) {
	return
	// stderr := os.Stderr
	// stderr.WriteString("\033[s\n")
	// for r := 0; r < len(grid); r++ {
	// 	for c := 0; c < len(grid[r]); c++ {
	// 		if grid[r][c] {
	// 			stderr.Write([]byte{'*'})
	// 		} else {
	// 			stderr.Write([]byte{' '})
	// 		}
	// 	}
	// 	stderr.Write([]byte{'\n'})
	// }
	// stderr.WriteString("\033[u")
	// time.Sleep(50 * time.Millisecond)
}

func findPath(grid [][]bool, trail []Point, trailLen int) bool {
	printGrid(grid)
	R := len(grid)
	C := len(grid[0])
	if trailLen == R*C {
		return true
	}
	candidatePoints := make([]Point, 0)
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] {
				continue
			}
			if trailLen == 0 {
				candidatePoints = append(candidatePoints, Point{r, c})
			} else {
				lastPoint := trail[trailLen-1]
				assert(grid[lastPoint.r][lastPoint.c])
				if lastPoint.r == r {
					break // skip this row
				}
				if lastPoint.c == c {
					continue
				}
				if lastPoint.r-lastPoint.c == r-c || lastPoint.r+lastPoint.c == r+c {
					continue
				}
				candidatePoints = append(candidatePoints, Point{r, c})
			}
		}
	}
	if len(candidatePoints) == 0 {
		return false
	} else {
		shufflePoints(candidatePoints)
		for _, cp := range candidatePoints {
			grid[cp.r][cp.c] = true
			trail[trailLen] = Point{cp.r, cp.c}
			ok := findPath(grid, trail, trailLen+1)
			if !ok {
				grid[cp.r][cp.c] = false
				continue
			}
			return true
		}
		return false
	}
}

// CJ's runner runs go 1.7.4

func shufflePoints(ps []Point) {
	for i := 0; i < len(ps)-1; i++ {
		j := rand.Intn(len(ps)-i) + i
		ps[i], ps[j] = ps[j], ps[i]
	}
}

func test(R, C int) {
	if R <= 2 && C <= 2 {
		output(false)
		return
	}
	grid := make([][]bool, R)
	for r := 0; r < R; r++ {
		grid[r] = make([]bool, C)
	}
	trail := make([]Point, R*C)
	ok := findPath(grid, trail, 0)
	printGrid(grid)
	output(ok)
	if ok {
		for _, r := range trail {
			fmt.Fprintf(stdout, "%d %d\n", r.r+1, r.c+1)
		}
	}
}

func output(possible bool) {
	if possible {
		stdout.WriteString("POSSIBLE\n")
	} else {
		stdout.WriteString("IMPOSSIBLE\n")
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
