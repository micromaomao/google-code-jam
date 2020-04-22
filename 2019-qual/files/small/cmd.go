/*
 THIS SOLUTION IS THE ONE I SUBMITTED DURING THE ROUND AND DID NOT PASS THE HIDDEN TEST!
*/

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		result := round()
		wrote := false
		for _, seg := range result {
			if seg.numBad == seg.length {
				for p := seg.start; p < seg.start+seg.length; p++ {
					if wrote {
						stdout.WriteByte(' ')
					}
					stdout.WriteString(strconv.Itoa(p))
					wrote = true
				}
			}
		}
		stdout.WriteByte('\n')
		stdout.Flush()
		var verdict int
		mustReadLineOfInts(&verdict)
		if verdict != 1 {
			return
		}
	}
}

type Seg struct {
	start           int
	length          int
	numBad          int
	test_interlaced bool
}

func (seg Seg) solved() bool {
	return seg.length == seg.numBad || seg.numBad == 0
}

func round() []Seg {
	var nWorkers, nBad, qLimit, qUsed int
	mustReadLineOfInts(&nWorkers, &nBad, &qLimit)
	var segments = make([]Seg, 0, 1)
	segments = append(segments, Seg{
		start:  0,
		length: nWorkers,
		numBad: nBad,
	})
	for {
		allSolved := true
		query := make([]int, 0, nWorkers)
		for segI := 0; segI < len(segments); segI++ {
			seg := &segments[segI]
			if seg.solved() {
				for i := 0; i < seg.length; i++ {
					query = append(query, 0)
				}
			} else {
				allSolved = false
				if seg.numBad != 1 {
					mid := seg.length / 2
					for i := 0; i < mid; i++ {
						query = append(query, 1)
					}
					for i := mid; i < seg.length; i++ {
						query = append(query, 0)
					}
				} else {
					// interlaced: 10101010
					next1 := true
					for i := 0; i < seg.length; i++ {
						if next1 {
							query = append(query, 1)
						} else {
							query = append(query, 0)
						}
						next1 = !next1
					}
					seg.test_interlaced = true
				}
			}
		}
		if allSolved {
			return segments
		}
		if qUsed >= qLimit {
			return segments // -.-
		}
		qUsed++
		for _, i := range query {
			if i == 0 {
				stdout.WriteByte('0')
			} else {
				stdout.WriteByte('1')
			}
		}
		stdout.WriteByte('\n')
		stdout.Flush()
		response := mustReadLine()
		if response == "-1" {
			return segments // Something's wrong.
		}
		var index int
		newSegments := make([]Seg, 0, len(segments)*2)
		for _, seg := range segments {
			relevantRes := response[index : index+seg.length-seg.numBad]
			index += len(relevantRes)
			if seg.solved() {
				newSegments = append(newSegments, seg)
			} else if !seg.test_interlaced {
				leftLen := seg.length / 2
				// 11110000
				var num1s int
				for _, r := range relevantRes {
					if r == '1' {
						num1s++
					} else {
						break
					}
				}
				segLeft := Seg{
					start:  seg.start,
					length: leftLen,
					numBad: leftLen - num1s,
				}
				segRight := Seg{
					start:  seg.start + leftLen,
					length: seg.length - leftLen,
					numBad: seg.numBad - segLeft.numBad,
				}
				newSegments = append(newSegments, segLeft, segRight)
			} else {
				next1 := true
				var skippedPos int
				for skippedPos = 0; skippedPos < len(relevantRes); skippedPos++ {
					if !next1 == (relevantRes[skippedPos] == '0') {
						next1 = !next1
						continue
					}
					break
				}
				// 1010101
				// 101 101
				// skippedPos = 3
				if skippedPos > 0 {
					newSegments = append(newSegments, Seg{
						start:  seg.start,
						length: skippedPos,
						numBad: 0,
					})
				}
				newSegments = append(newSegments, Seg{
					start:  seg.start + skippedPos,
					length: 1,
					numBad: 1,
				})
				if skippedPos < seg.length-1 {
					newSegments = append(newSegments, Seg{
						start:  seg.start + skippedPos + 1,
						length: seg.length - skippedPos - 1,
					})
				}
			}
		}
		segments = newSegments
	}
}

func writeAndFlush(str string) {
	stdout.WriteString(str)
	stdout.Flush()
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
