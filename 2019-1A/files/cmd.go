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

type SuffixTree struct {
	strsEndsIn []int
	subTree    map[byte]*SuffixTree
}

func newSuffixTree() *SuffixTree {
	return &SuffixTree{strsEndsIn: make([]int, 0), subTree: make(map[byte]*SuffixTree)}
}

func (t *SuffixTree) addInTree(suffix string, stridx int) {
	leaf := t
	for i := 0; i < len(suffix); i++ {
		char := suffix[i]
		if _, ok := leaf.subTree[char]; !ok {
			leaf.subTree[char] = newSuffixTree()
		}
		leaf = leaf.subTree[char]
	}
	leaf.strsEndsIn = append(leaf.strsEndsIn, stridx)
}

func (t *SuffixTree) find(suffix string) []int {
	leaf := t
	for i := 0; i < len(suffix); i++ {
		char := suffix[i]
		if _, ok := leaf.subTree[char]; !ok {
			return nil
		}
		leaf = leaf.subTree[char]
	}
	return leaf.strsEndsIn
}

func test() {
	var N int
	mustReadLineOfInts(&N)
	words := make([]string, 0, N)
	// t := newSuffixTree()
	maxLength := 0
	for i := 0; i < N; i++ {
		word := mustReadLine()
		words = append(words, word)
		if maxLength < len(word) {
			maxLength = len(word)
		}
		// stridx := len(words)
		// for i := 0; i < len(word); i++ {
		// 	t.addInTree(word[i:], stridx)
		// }
	}
	maxSuffixLength := make([]int, N)
	for i := 0; i < N; i++ {
		maxSuffixLength[i] = len(words[i])
	}
	pairedWordsCount := 0
	// for thisWord := 0; thisWord < N; thisWord++ {
	// 	if !wordsRemaining[thisWord] {
	// 		continue
	// 	}
	// 	word := words[thisWord]
	// 	for suf := 0; suf < len(word); suf++ {
	// 		_strs := t.find(word[suf:])
	// 		strsWithThisSuf := make([]int, 0, len(_strs))
	// 		for _, stridx := range _strs {
	// 			if wordsRemaining[stridx] {
	// 				strsWithThisSuf = append(strsWithThisSuf, stridx)
	// 				if len(strsWithThisSuf) > 2 {
	// 					break
	// 				}
	// 			}
	// 		}
	// 		if len(strsWithThisSuf) > 2 {
	// 			wordsRemaining[thisWord] = false
	// 			break
	// 		}
	// 		if len(strsWithThisSuf) == 2 {
	// 			otherWord := strsWithThisSuf[0]
	// 			if otherWord == thisWord {
	// 				otherWord = strsWithThisSuf[1]
	// 			}
	// 			pairedWordsCount += 2
	// 			wordsRemaining[otherWord] = false
	// 			wordsRemaining[thisWord] = false
	// 			break
	// 		}
	// 		assert(len(strsWithThisSuf) == 1)
	// 	}
	// }

	for suffixLength := maxLength; suffixLength >= 1; suffixLength-- {
		// debug(fmt.Sprintf("suffixLength = %v", suffixLength))
		t := newSuffixTree()
		for thisWord := 0; thisWord < len(words); thisWord++ {
			if maxSuffixLength[thisWord] < suffixLength {
				continue
			}
			word := words[thisWord]
			if len(word) < suffixLength {
				continue
			}
			thisSuffix := word[len(word)-suffixLength:]
			// debug(fmt.Sprintf("Adding %v", thisSuffix))
			t.addInTree(thisSuffix, thisWord)
		}
		for thisWord := 0; thisWord < len(words); thisWord++ {
			if maxSuffixLength[thisWord] < suffixLength {
				continue
			}
			word := words[thisWord]
			if len(word) < suffixLength {
				continue
			}
			thisSuffix := word[len(word)-suffixLength:]
			// debug(thisSuffix)
			_foundWords := t.find(thisSuffix)
			foundWords := make([]int, 0, len(_foundWords))
			for _, w := range _foundWords {
				if maxSuffixLength[w] >= suffixLength {
					foundWords = append(foundWords, w)
				}
			}
			if len(foundWords) >= 2 {
				otherWord := foundWords[0]
				if otherWord == thisWord {
					otherWord = foundWords[1]
				}
				maxSuffixLength[thisWord] = 0
				maxSuffixLength[otherWord] = 0
				pairedWordsCount += 2
				// debug(fmt.Sprintf("Pair: %v and %v. Found %v other words", word, words[otherWord], len(foundWords)-1))
				for i := 2; i < len(foundWords); i++ {
					maxSuffixLength[foundWords[i]] = suffixLength - 1
					// debug(fmt.Sprintf("Decreasing suffixLength for %v to %v", words[foundWords[i]], maxSuffixLength[foundWords[i]]))
				}
				continue
			}
		}
	}
	fmt.Fprintf(stdout, "%d\n", pairedWordsCount)
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
