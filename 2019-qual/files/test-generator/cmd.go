package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
)

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

func main() {
	fin, err := os.OpenFile("generated.in", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	fout, err := os.OpenFile("generated.out", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	in := bufio.NewWriter(fin)
	out := bufio.NewWriter(fout)
	t := 1000000
	N := 10000
	in.WriteString(strconv.Itoa(t))
	in.WriteByte('\n')
	for i := 0; i < t; i++ {
		in.WriteString(strconv.Itoa(N))
		in.WriteByte(' ')
		L := rand.Intn(100) + 25
		in.WriteString(strconv.Itoa(L))
		in.WriteByte('\n')
		out.WriteString("Case #")
		out.WriteString(strconv.Itoa(i + 1))
		out.WriteString(": ")
		primesList := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
		msg := make([]rune, 0, L+1)
		for i := 0; i < len(primesList); i++ {
			msg = append(msg, 'A'+rune(i))
		}
		for i := len(msg); i < L+1; i++ {
			idx := rand.Intn(26)
			msg = append(msg, 'A'+rune(idx))
		}
		rand.Shuffle(len(msg), func(i, j int) {
			msg[i], msg[j] = msg[j], msg[i]
		})
		out.WriteString(string(msg))
		out.WriteByte('\n')
		for i := 0; i < len(msg)-1; i++ {
			prod := primesList[byte(msg[i])-'A'] * primesList[byte(msg[i+1])-'A']
			if i != 0 {
				in.WriteByte(' ')
			}
			in.WriteString(strconv.Itoa(prod))
		}
		in.WriteByte('\n')
	}
	out.Flush()
	in.Flush()
}
