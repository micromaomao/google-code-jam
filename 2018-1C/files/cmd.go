// Lesson: the complexity of a tree traversal won't blow up if only valid inputs are traversed. Don't be afraid to
// traverse a word tree.
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

type Tree struct {
	charId     int
	subTrees   map[int]*Tree
	idsNotUsed map[int]bool
}

func newTree(charId int, total int) *Tree {
	idsNotUsed := make(map[int]bool)
	for i := 0; i < total; i++ {
		idsNotUsed[i] = true
	}
	return &Tree{charId: charId, subTrees: make(map[int]*Tree), idsNotUsed: idsNotUsed}
}

func (t *Tree) find(charId int, nextTotals int) *Tree {
	subt := t.subTrees[charId]
	if subt == nil {
		t.subTrees[charId] = newTree(charId, nextTotals)
		delete(t.idsNotUsed, charId)
	}
	return t.subTrees[charId]
}

func test() {
	var N, L int
	mustReadLineOfInts(&N, &L)
	words := make([]string, N)
	for i := 0; i < N; i++ {
		words[i] = mustReadLine()
		assert(len(words[i]) == L)
	}
	translationTables := make([]map[byte]int, L)
	revTranslationTables := make([]map[int]byte, L)
	totals := make([]int, L)
	for i := 0; i < L; i++ {
		cnt := 0
		tt := make(map[byte]int)
		translationTables[i] = tt
		revTt := make(map[int]byte)
		revTranslationTables[i] = revTt
		for j := 0; j < N; j++ {
			char := words[j][i]
			_, existed := tt[char]
			if !existed {
				tt[char] = cnt
				revTt[cnt] = char
				cnt++
			}
		}
		totals[i] = cnt
	}
	t := newTree(-1, totals[0])
	for i := 0; i < N; i++ {
		word := words[i]
		currentT := t
		for j := 0; j < L; j++ {
			char := word[j]
			nextTotal := 0
			if j < L-1 {
				nextTotal = totals[j+1]
			}
			currentT = currentT.find(translationTables[j][char], nextTotal)
		}
	}

	var trav func(t *Tree, buf []int) bool
	trav = func(t *Tree, buf []int) bool {
		if len(t.idsNotUsed) > 0 {
			var oneNotUsedId int
			for id, b := range t.idsNotUsed {
				assert(b)
				oneNotUsedId = id
				break
			}
			buf[0] = oneNotUsedId
			for i := 1; i < len(buf); i++ {
				buf[i] = 0
			}
			return true
		} else {
			if len(buf) == 1 {
				return false
			}
			assert(len(buf) > 1)
			for sid, t := range t.subTrees {
				buf[0] = sid
				if trav(t, buf[1:]) {
					return true
				}
			}
			return false
		}
	}

	wordIds := make([]int, L)
	if trav(t, wordIds) {
		word := make([]byte, L)
		for i := 0; i < L; i++ {
			word[i] = revTranslationTables[i][wordIds[i]]
		}
		stdout.Write(word)
		stdout.WriteByte('\n')
	} else {
		stdout.WriteString("-\n")
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
