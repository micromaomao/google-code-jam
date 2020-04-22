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
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		ls := test()
		if ls == nil {
			stdout.WriteString("IMPOSSIBLE\n")
		} else {
			stdout.Write(ls)
			stdout.WriteByte('\n')
		}
	}
}

func test() []byte {
	var N, R, O, Y, G, B, V int
	mustReadLineOfInts(&N, &R, &O, &Y, &G, &B, &V)
	preList := make([][]byte, 0, 3)
	nR := R
	nY := Y
	nB := B
	if O > 0 {
		if B == O && N == O*2 {
			rep := make([]byte, N)
			for i := 0; i < N; i += 2 {
				rep[i] = 'B'
				rep[i+1] = 'O'
			}
			return rep
		}
		if B < O+1 {
			return nil
		}
		B -= O + 1
		ll := make([]byte, O*2+1)
		ll[0] = 'B'
		for i := 1; i < len(ll); i += 2 {
			ll[i] = 'O'
		}
		for i := 2; i < len(ll); i += 2 {
			ll[i] = 'B'
		}
		O = 0
		preList = append(preList, ll)
		nB = B + 1
	}
	if G > 0 {
		if R == G && N == G*2 {
			rep := make([]byte, N)
			for i := 0; i < N; i += 2 {
				rep[i] = 'R'
				rep[i+1] = 'G'
			}
			return rep
		}
		if R < G+1 {
			return nil
		}
		R -= G + 1
		ll := make([]byte, G*2+1)
		ll[0] = 'R'
		for i := 1; i < len(ll); i += 2 {
			ll[i] = 'G'
		}
		for i := 2; i < len(ll); i += 2 {
			ll[i] = 'R'
		}
		G = 0
		preList = append(preList, ll)
		nR = R + 1
	}
	if V > 0 {
		if Y == V && N == V*2 {
			rep := make([]byte, N)
			for i := 0; i < N; i += 2 {
				rep[i] = 'Y'
				rep[i+1] = 'V'
			}
			return rep
		}
		if Y < V+1 {
			return nil
		}
		Y -= V + 1
		ll := make([]byte, V*2+1)
		ll[0] = 'Y'
		for i := 1; i < len(ll); i += 2 {
			ll[i] = 'V'
		}
		for i := 2; i < len(ll); i += 2 {
			ll[i] = 'Y'
		}
		V = 0
		preList = append(preList, ll)
		nY = Y + 1
	}
	assert(len(preList) <= 3)
	preListMap := make(map[byte][]byte)
	for _, ll := range preList {
		preListMap[ll[0]] = ll
	}
	debug(fmt.Sprintf("%v", preListMap))
	baseList := sortOutSimpleCase(nR, nY, nB)
	if baseList == nil {
		return nil
	} else {
		newList := make([]byte, 0, N)
		for _, baseLet := range baseList {
			if pl, exist := preListMap[baseLet]; exist {
				for _, letter := range pl {
					newList = append(newList, letter)
				}
				delete(preListMap, baseLet)
			} else {
				newList = append(newList, baseLet)
			}
		}
		return newList
	}
}

type LetterAndNumber struct {
	letter byte
	number int
}

func sortOutSimpleCase(R, Y, B int) []byte {
	colors := make([]LetterAndNumber, 3)
	colors[0] = LetterAndNumber{'R', R}
	colors[1] = LetterAndNumber{'Y', Y}
	colors[2] = LetterAndNumber{'B', B}
	sort.Slice(colors, func(i, j int) bool {
		return colors[i].number > colors[j].number
	})
	dbgStr := make([]string, 0)
	for _, c := range colors {
		dbgStr = append(dbgStr, fmt.Sprintf("%v %v", string(c.letter), c.number))
	}
	debug(strings.Join(dbgStr, ", "))
	if colors[0].number == 0 {
		return make([]byte, 0)
	}
	if colors[1].number == 0 {
		return nil
	}
	list := make([]byte, R+Y+B)
	for a := 0; a < len(list); a++ {
		list[a] = '_'
	}
	var i int
	for i = 0; i < len(list)-1; i += 2 {
		list[i] = colors[0].letter
		colors[0].number--
		if colors[0].number == 0 {
			break
		}
	}
	if colors[0].number > 0 {
		return nil
	}
	i += 2
	if i >= len(list) {
		i = 1
	} else {
		for ; i < len(list); i += 2 {
			if colors[1].number == 0 {
				panic("!")
			}
			list[i] = colors[1].letter
			colors[1].number--
		}
	}
	i = 1
	for ; i < len(list); i += 2 {
		if colors[1].number > 0 {
			list[i] = colors[1].letter
			colors[1].number--
		} else {
			assert(colors[2].number > 0)
			list[i] = colors[2].letter
			colors[2].number--
		}
	}
	for _, l := range list {
		assert(l != '_')
	}
	return list
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
