package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	fin, err := os.OpenFile("generated.in", os.O_WRONLY|os.O_CREATE, 0666)
	finHelp, err := os.OpenFile("generated.in.wcomment", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	in := bufio.NewWriter(fin)
	inHelp := bufio.NewWriter(finHelp)
	t := 10000
	fmt.Fprintf(in, "%d\n", t)
	fmt.Fprintf(inHelp, "%d\n", t)
	type Cookie struct {
		w int
		h int
	}
	for i := 0; i < t; i++ {
		var P int = 0
		var maxP float64 = 0
		N := rand.Intn(10) + 1
		cookies := make([]Cookie, 0, N)
		for c := 0; c < N; c++ {
			cookie := Cookie{
				w: rand.Intn(249) + 1,
				h: rand.Intn(249) + 1,
			}
			cookies = append(cookies, cookie)
			P += 2 * (cookie.w + cookie.h)
			maxP += 2 * (float64(cookie.w) + float64(cookie.h) + math.Sqrt(float64(cookie.w*cookie.w)+float64(cookie.h*cookie.h)))
		}
		if rand.Intn(2) == 1 {
			P += rand.Intn(int(maxP-float64(P)) * 2)
		}
		fmt.Fprintf(in, "%d %d\n", N, P)
		fmt.Fprintf(inHelp, "%d %d // Start case %d\n", N, P, i+1)
		for _, cookie := range cookies {
			fmt.Fprintf(in, "%d %d\n", cookie.w, cookie.h)
			fmt.Fprintf(inHelp, "%d %d\n", cookie.w, cookie.h)
		}
	}
	in.Flush()
	inHelp.Flush()
}
