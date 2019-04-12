package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		stdout.WriteString("Case #")
		stdout.WriteString(strconv.Itoa(i + 1))
		stdout.WriteString(": ")
		test()
	}
}

type PrimeList []*big.Int

func (pl *PrimeList) Len() int {
	return len(*pl)
}

func (pl *PrimeList) Less(i int, j int) bool {
	return (*pl)[i].Cmp((*pl)[j]) < 0
}

func (pl *PrimeList) Swap(i int, j int) {
	tmp := (*pl)[i]
	(*pl)[i] = (*pl)[j]
	(*pl)[j] = tmp
}

func (pl PrimeList) SearchFor(i *big.Int) int {
	res := sort.Search(len(pl), func(idx int) bool {
		return pl[idx].Cmp(i) >= 0
	})
	if pl[res].Cmp(i) != 0 {
		debug(fmt.Sprintf("Warning: searched for %v, found index = %v, pl[index] = %v which does not equal the query.", i.Text(10), res, pl[res].Text(10)))
	}
	return res
}

func test() {
	zero := big.NewInt(0)
	var L int
	line := strings.Split(mustReadLine(), " ")
	// N, success := big.NewInt(0).SetString(line[0], 10)
	// if !success {
	// 	panic("!")
	// }
	L = mustAtoi(line[1])
	cipherText := make([]*big.Int, 0, L)
	line = strings.Split(mustReadLine(), " ")
	if len(line) != L {
		panic("length mismatch.")
	}
	primesFound := make(PrimeList, 0)
	for _, numStr := range line {
		bn, success := big.NewInt(0).SetString(numStr, 10)
		if !success {
			panic("!")
		}
		cipherText = append(cipherText, bn)
	}
	problematicProducts := make([]*big.Int, 0)
	firstNonProblematicCipherIndex := -1
	lastNonProblematicCipherIndex := -1
	for k := 0; k < L-1; k++ {
		if cipherText[k].Cmp(cipherText[k+1]) == 0 {
			problematicProducts = append(problematicProducts, cipherText[k])
			continue
		}
		if firstNonProblematicCipherIndex == -1 {
			firstNonProblematicCipherIndex = k
		}
		lastNonProblematicCipherIndex = k + 1
		p := big.NewInt(0).GCD(nil, nil, cipherText[k], cipherText[k+1])
		primesFound = append(primesFound, p)
	}
	firstPrimeFound := primesFound[0]
	lastPrimeFound := primesFound[len(primesFound)-1]
	var firstPrime *big.Int = nil
	if firstNonProblematicCipherIndex%2 == 0 {
		firstPrime = big.NewInt(0).Div(cipherText[firstNonProblematicCipherIndex], firstPrimeFound)
	} else {
		firstPrime = firstPrimeFound
	}
	lastPrime := big.NewInt(0).Div(cipherText[lastNonProblematicCipherIndex], lastPrimeFound)
	primesFound = append(primesFound, firstPrime, lastPrime)
	buffer := big.NewInt(0)
	buffer2 := big.NewInt(0)
	for _, prod := range problematicProducts {
		for _, primeToTry := range primesFound {
			buffer, m := buffer.DivMod(prod, primeToTry, buffer2)
			if m.Cmp(zero) == 0 {
				q := big.NewInt(0).Set(buffer)
				primesFound = append(primesFound, q)
				break
			}
		}
	}
	sort.Sort(&primesFound)
	last := zero
	nPrimesFound := make(PrimeList, 0)
	for _, p := range primesFound {
		if p.Cmp(last) != 0 {
			nPrimesFound = append(nPrimesFound, p)
			last = p
		}
	}
	primesFound = nPrimesFound
	//debug(fmt.Sprintf("%v", primesFound))
	if len(primesFound) != 26 {
		debug("Warning: not exactly 26 primes! Got " + strconv.Itoa(len(primesFound)))
	}
	lastDecrypted := primesFound.SearchFor(firstPrime)
	stdout.WriteByte('A' + byte(lastDecrypted))
	for _, c := range cipherText {
		buffer := buffer.Div(c, primesFound[lastDecrypted])
		t := primesFound.SearchFor(buffer)
		lastDecrypted = t
		stdout.WriteByte('A' + byte(t))
	}
	stdout.WriteByte('\n')
}

func encrypt(primes []int, msg string) []int {
	nums := make([]int, 0, len(msg))
	for _, char := range msg {
		if char-'A' > 26 {
			continue
		}
		nums = append(nums, primes[char-'A'])
	}
	c := make([]int, 0, len(msg)-1)
	for i := 0; i < len(nums)-1; i++ {
		c = append(c, nums[i]*nums[i+1])
	}
	return c
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
