package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

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
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Expected name of program as the only argument, got %d arguments instead.", len(os.Args))
		os.Exit(1)
	}
	stdout := bufio.NewWriter(os.Stdout)
	cmd := exec.Command(os.Args[1])
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	rawout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
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
				p := recover()
				if p != nil {
					res = false
				}
			}()
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
	stdout.Flush()
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

func test(in *WrappedWriter, out *WrappedReader) bool {
	line := out.ReadLine()
	assert(line == "1")
	return true
}
