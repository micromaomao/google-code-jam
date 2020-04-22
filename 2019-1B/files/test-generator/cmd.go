package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	rawIn, err := os.OpenFile("generated.in", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	in := bufio.NewWriter(rawIn)
	defer rawIn.Close()
	defer in.Flush()
	const t int = 100
	fmt.Fprintf(in, "%d\n", t)
	for i := 0; i < t; i++ {
		var N, K int
		N = rand.Intn(10) + 1
		K = rand.Intn(5)
		fmt.Fprintf(in, "%d %d\n", N, K)
		for i := 0; i < 2; i++ {
			for i := 0; i < N; i++ {
				if i != 0 {
					in.WriteByte(' ')
				}
				in.WriteString(strconv.Itoa(rand.Intn(15)))
			}
			in.WriteByte('\n')
		}
	}
}
