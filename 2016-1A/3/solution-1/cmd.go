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

type Node struct {
	childId   int
	bff       *Node
	isBffOf   []*Node
	distance  int
	component *Component
	isInCycle bool
}

type Component struct {
	entry      *Node
	cycleSize  int
	isPartOf   *Component
	cycleEntry *Node
}

func test() {
	stdout.Flush()
	os.Stderr.Write([]byte{'\n'})
	var N int
	mustReadLineOfInts(&N)
	bffs := mustReadLineOfIntsIntoArray()
	assert(len(bffs) == N)
	for i := 0; i < N; i++ {
		bffs[i]--
		assert(bffs[i] >= 0 && bffs[i] < N && bffs[i] != i)
	}
	childsLeft := make(map[int]bool)
	for i := 0; i < N; i++ {
		childsLeft[i] = true
	}
	childNodes := make([]*Node, N)
	components := make([]Component, 0)
	for len(childsLeft) > 0 {
		var childId int
		for childId, _ = range childsLeft {
			break
		}
		components = append(components, Component{})
		cp := &components[len(components)-1]
		entry := buildChild(cp, bffs, childNodes, childsLeft, childId, 0)
		cp.entry = entry
		if cp.isPartOf == nil {
			debug(fmt.Sprintf("Component with child %d has cycle size %d starting %v", cp.entry.childId+1, cp.cycleSize, cp.cycleEntry.childId+1))
		}
	}
	maxCycleSize := 0
	twoCyclesSzie := 0
	for _, c := range components {
		if c.isPartOf != nil {
			continue
		}
		if c.cycleSize > 2 {
			if maxCycleSize < c.cycleSize {
				maxCycleSize = c.cycleSize
			}
		} else {
			assert(c.cycleSize == 2)
			entry := c.cycleEntry
			assert(c.cycleEntry != nil)
			current := entry
			ct := 0
			for {
				ct++
				current.isInCycle = true
				current = current.bff
				if current == entry {
					break
				}
			}
			assert(ct == 2)
			ct = 0
			tl := 0
			for {
				ct++
				tl += tailSize(current)
				current = current.bff
				if current == entry {
					break
				}
			}
			assert(ct == 2)
			ct = 0
			cSize := 2 + tl
			twoCyclesSzie += cSize
		}
	}
	maxCSize := maxCycleSize
	if twoCyclesSzie > maxCSize {
		maxCSize = twoCyclesSzie
	}
	fmt.Fprintf(stdout, "%d\n", maxCSize)
}

func buildChild(component *Component, bffs []int, childNodes []*Node, childsLeft map[int]bool, childId int, distance int) *Node {
	if childNodes[childId] != nil {
		_, exist := childsLeft[childId]
		assert(!exist)
		cn := childNodes[childId]
		if cn.component == component {
			assert(component.cycleSize == 0)
			component.cycleSize = distance - cn.distance
			component.cycleEntry = cn
			return cn
		} else {
			component.isPartOf = cn.component
			// We are not in a cycle
			return cn
		}
	}
	_, exist := childsLeft[childId]
	assert(exist)
	delete(childsLeft, childId)
	childNodes[childId] = &Node{
		childId:   childId,
		bff:       nil,
		isBffOf:   make([]*Node, 0),
		distance:  distance,
		component: component,
		isInCycle: false,
	}
	bff := buildChild(component, bffs, childNodes, childsLeft, bffs[childId], distance+1)
	childNodes[childId].bff = bff
	bff.isBffOf = append(bff.isBffOf, childNodes[childId])
	if bff.component != component {
		childNodes[childId].component = bff.component
		assert(bff.component == component.isPartOf)
	}
	return childNodes[childId]
}

func tailSize(node *Node) int {
	maxTailSize := 0
	for _, t := range node.isBffOf {
		assert(t.component == node.component)
		if t.isInCycle {
			continue
		}
		ts := tailSize(t) + 1
		if maxTailSize < ts {
			maxTailSize = ts
		}
	}
	return maxTailSize
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
