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
		stdout.WriteString(fmt.Sprintf("Case #%d: ", i+1))
		test()
	}
}

type Cookie struct {
	w              int
	h              int
	pAdditionalMin float64
	pAdditionalMax float64
}

const eplison float64 = 0.000001

type Interval struct {
	min float64
	max float64
}

func test() {
	var N, P int
	mustReadLineOfInts(&N, &P)
	cookies := make([]Cookie, 0, N)
	for i := 0; i < N; i++ {
		var w, h int
		mustReadLineOfInts(&w, &h)
		cookies = append(cookies, Cookie{
			w: w, h: h,
		})
	}
	var baseP float64
	for i := 0; i < len(cookies); i++ {
		cookie := &cookies[i]
		if cookie.w < cookie.h {
			cookie.w, cookie.h = cookie.h, cookie.w
		}
		cookieP := 2 * (cookie.w + cookie.h)
		cookie.pAdditionalMin = float64(2*(cookie.w+2*cookie.h) - cookieP)
		cookie.pAdditionalMax = 2*(float64(cookie.w+cookie.h)+float64(math.Sqrt(float64(cookie.w*cookie.w+cookie.h*cookie.h)))) - float64(cookieP)
		baseP += float64(cookieP)
	}
	addIntervals := make(IntervalArray, 0)
	addIntervals = append(addIntervals, Interval{0, 0}) // In the loop below, each cookie would have a chance for its addMin, addMax to be added to this 0, 0 interval, representing the case where only that cookie is cutted.
	maxMin := float64(P) - baseP
	if maxMin < eplison {
		fmt.Fprintf(stdout, "%.15f\n", float64(P))
		return
	}
	for i := 0; i < len(cookies); i++ {
		cookie := &cookies[i]
		// Either don't cut this cookie, in which case intervals are untouched, or...
		for _, interval := range addIntervals {
			addIntervals = append(addIntervals, Interval{interval.min + cookie.pAdditionalMin, interval.max + cookie.pAdditionalMax})
		}
		mergeIntervals(&addIntervals, maxMin)
	}
	lastInterval := addIntervals[len(addIntervals)-1]
	maxP := lastInterval.max + baseP
	if maxP > float64(P) {
		fmt.Fprintf(stdout, "%.15f\n", float64(P))
	} else {
		fmt.Fprintf(stdout, "%.15f\n", maxP)
	}
}

type IntervalArray []Interval

func (its IntervalArray) Len() int {
	return len(its)
}

func (its IntervalArray) Less(i int, j int) bool {
	return its[i].min < its[j].min
}

func (its IntervalArray) Swap(i int, j int) {
	its[i], its[j] = its[j], its[i]
}

func mergeIntervals(its *IntervalArray, maxMin float64) {
	if len(*its) == 0 {
		return
	}
	sort.Sort(*its)
	if (*its)[0].min > maxMin {
		*its = make(IntervalArray, 0)
		return
	}
	newArr := make(IntervalArray, 0)
	currentStarting := (*its)[0]
	for i := 1; i < len(*its); i++ {
		it := (*its)[i]
		if it.min-eplison > currentStarting.max {
			if it.min > maxMin {
				break
			}
			newArr = append(newArr, currentStarting)
			currentStarting = it
		} else if it.max > currentStarting.max {
			currentStarting.max = it.max
		}
	}
	newArr = append(newArr, currentStarting)
	*its = newArr
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
