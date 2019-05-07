package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func test(in *WrappedWriter, out *WrappedReader) bool {
	A := rand.Intn(7) + 1
	fmt.Fprintf(in, "%d\n", A)
	advPrograms := make([][]byte, A)
	for i := 0; i < A; i++ {
		prog := make([]byte, rand.Intn(5)+1)
		options := []byte{'R', 'P', 'S'}
		for j := 0; j < len(prog); j++ {
			prog[j] = options[rand.Intn(len(options))]
		}
		advPrograms[i] = prog
		in.Write(prog)
		in.Write([]byte{'\n'})
	}
	response := strings.Split(out.ReadLine(), ": ")
	assert(len(response) == 2)
	prog := response[1]
	assert(len(prog) <= 500)
	if prog == "IMPOSSIBLE" {
		// TODO
		return true
	} else {
	a:
		for _, oppProg := range advPrograms {
			maxLen := len(oppProg)
			if len(prog) > maxLen {
				maxLen = len(prog)
			}
			for pos := 0; pos < maxLen; pos++ {
				myAct := prog[pos%len(prog)]
				oppAct := oppProg[pos%len(oppProg)]
				res := round(myAct, oppAct)
				if res < 0 {
					return false
				}
				if res > 0 {
					continue a
				}
			}
			return false
		}
		return true
	}
}

func round(my, other byte) int8 {
	switch my {
	case 'R':
		if other == 'P' {
			return -1
		} else if other == 'S' {
			return 1
		}
	case 'P':
		if other == 'S' {
			return -1
		} else if other == 'R' {
			return 1
		}
	case 'S':
		if other == 'R' {
			return -1
		} else if other == 'P' {
			return 1
		}
	default:
		panic("!")
	}
	return 0
}

type WrappedReader struct {
	r      *bufio.Reader
	stdout io.Writer
}
type WrappedWriter struct {
	w         io.Writer
	stdout    io.Writer
	outBuffer *bufio.Writer
}

func (wr *WrappedReader) ReadLine() string {
	line, err := wr.r.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\n")
	fmt.Fprintf(wr.stdout, "< %s\n", line)
	return line
}

func (ww *WrappedWriter) Write(p []byte) (n int, err error) {
	n, err = ww.w.Write(p)
	if err != nil {
		if err == io.ErrClosedPipe {
			err = nil
			return
		}
		panic(err)
	}
	if ww.outBuffer == nil {
		ww.outBuffer = bufio.NewWriterSize(ww.stdout, 10000)
		ww.outBuffer.WriteString("> ")
	}
	for _, b := range p[:n] {
		if b == '\n' {
			ww.outBuffer.WriteByte('\n')
			ww.outBuffer.Flush()
			ww.outBuffer.WriteString("> ")
		} else {
			ww.outBuffer.WriteByte(b)
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Expected name of program as the only argument, got %d arguments instead.", len(os.Args))
		os.Exit(1)
	}
	// stdout := bufio.NewWriter(os.Stdout)
	stdout := os.Stdout
	cmd := exec.Command(os.Args[1])
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	rawout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Stderr = os.Stderr
	out := bufio.NewReader(rawout)
	cmd.Start()
	t := 1
	success := true
	ww := WrappedWriter{w: in, stdout: stdout}
	wr := WrappedReader{r: out, stdout: stdout}
	fmt.Fprintf(&ww, "%d\n", t)
	for i := 0; i < t; i++ {
		var res bool
		(func() {
			defer func() {
				// stdout.Flush()
				p := recover()
				if p != nil {
					res = false
				}
			}()
			// stdout.Flush()
			res = test(&ww, &wr)
		})()
		if !res || !success {
			success = false
			break
		}
	}
	in.Close()
	if ww.outBuffer.Buffered() > 2 {
		ww.outBuffer.Flush()
	}
	if !success {
		fmt.Fprintf(stdout, "Errored.\n")
	}
	// stdout.Flush()
	cmd.Process.Signal(syscall.SIGINT)
	cmd.Process.Wait()
	if !success {
		os.Exit(1)
	}
}

func assert(b bool) {
	if !b {
		panic("Assertion failed.")
	}
}
