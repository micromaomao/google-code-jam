package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		if i == 27 {
			stdout.Flush()
		}
		fmt.Fprintf(stdout, "%d\n", test())
	}
}

type Package struct {
	qty      int
	canDoMin int
	canDoMax int
	circled  int
}

func (p *Package) calc(req int) {
	p.canDoMax = int(math.Floor(float64(p.qty) / (0.9 * float64(req))))
	p.canDoMin = int(math.Ceil(float64(p.qty) / (1.1 * float64(req))))
	if req*(p.canDoMax+1)*9/10 <= p.qty {
		p.canDoMax++
	}
	if req*(p.canDoMin-1)*11/10 >= p.qty {
		p.canDoMin++
	}
	if p.canDoMax == 0 {
		p.canDoMin = 0
		p.canDoMax = -1
	}
}

type Packages []Package

func (p Packages) Len() int {
	return len(p)
}

func (p Packages) Less(i int, j int) bool {
	return p[i].qty < p[j].qty
}

func (p Packages) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func test() int {
	var N, P int
	mustReadLineOfInts(&N, &P)
	reqs := mustReadLineOfIntsIntoArray()
	assert(len(reqs) == N)
	packages := make([]Packages, N)
	for i := 0; i < N; i++ {
		pkgs := make(Packages, P)
		q := mustReadLineOfIntsIntoArray()
		assert(len(q) == P)
		for j := 0; j < P; j++ {
			pkgs[j].qty = q[j]
			pkgs[j].calc(reqs[i])
		}
		sort.Sort(pkgs)
		newPkgs := make(Packages, 0, P)
		for _, pkg := range pkgs {
			if pkg.canDoMin <= pkg.canDoMax {
				newPkgs = append(newPkgs, pkg)
			}
		}
		packages[i] = newPkgs
	}
	for _, pkgs := range packages {
		if len(pkgs) == 0 {
			return 0
		}
	}
	debug(fmt.Sprintf("%v", packages))
	if N == 1 {
		return len(packages[0])
	}
	for i := 1; i < N; i++ {
		prevRow := packages[i-1]
		thisRow := packages[i]
		for {
			if len(prevRow) == 0 || len(thisRow) == 0 {
				break
			}
			if thisRow[0].canDoMax < prevRow[0].canDoMin {
				thisRow = thisRow[1:]
				continue
			}
			if thisRow[0].canDoMin > prevRow[0].canDoMax {
				prevRow = prevRow[1:]
				continue
			}
			if prevRow[0].circled != 0 && thisRow[0].canDoMax < prevRow[0].circled {
				thisRow = thisRow[1:]
				continue
			}
			if i != 1 {
				if prevRow[0].circled != 0 {
					assert(prevRow[0].canDoMin <= prevRow[0].circled)
					prevRow[0].canDoMin = prevRow[0].circled
					assert(prevRow[0].canDoMax >= prevRow[0].canDoMin)
				} else {
					prevRow = prevRow[1:]
					continue
				}
			}
			if thisRow[0].canDoMin < prevRow[0].canDoMin {
				prevRow[0].circled = prevRow[0].canDoMin
				thisRow[0].circled = prevRow[0].canDoMin
			} else {
				thisRow[0].circled = thisRow[0].canDoMin
				prevRow[0].circled = thisRow[0].canDoMin
			}
			assert(thisRow[0].circled >= thisRow[0].canDoMin)
			assert(thisRow[0].circled <= thisRow[0].canDoMax)
			assert(prevRow[0].circled >= prevRow[0].canDoMin)
			assert(prevRow[0].circled <= prevRow[0].canDoMax)
			thisRow = thisRow[1:]
			prevRow = prevRow[1:]
		}
	}
	lastRow := packages[N-1]
	debug(fmt.Sprintf("%v", packages))
	kits := 0
	for _, p := range lastRow {
		if p.circled != 0 {
			kits++
		}
	}
	return kits
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
