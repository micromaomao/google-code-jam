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
		n := mustReadLine()
		stdout.WriteString("Case #")
		stdout.WriteString(strconv.Itoa(i + 1))
		stdout.WriteString(": ")
		bNum := make([]bool, len(n))
		for i := 0; i < len(n); i++ {
			if n[i] == '4' {
				bNum[i] = true
				stdout.WriteByte('3')
			} else {
				stdout.WriteByte(byte(n[i]))
			}
		}
		stdout.WriteString(" ")
		initialDigitWritten := false
		for i := 0; i < len(bNum); i++ {
			if !bNum[i] {
				if initialDigitWritten {
					stdout.WriteByte('0')
				}
			} else {
				initialDigitWritten = true
				stdout.WriteByte('1')
			}
		}
		if !initialDigitWritten {
			stdout.WriteByte('0')
		}
		stdout.WriteByte('\n')
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
