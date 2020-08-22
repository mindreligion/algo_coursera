package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const T = 666

func main() {
	f, err := os.OpenFile("./2_sum_test.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	var nn []int
	for s.Scan() {
		err := s.Err()
		if err != nil {
			panic(err)
		}
		line := s.Text()
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(line)
			panic(err)
		}

		nn = append(nn, n)
	}
	start := time.Now()
	pairs := t2s(nn)
	t := time.Now().Sub(start)
	fmt.Println(pairs)
	fmt.Println()
	fmt.Println(t / time.Millisecond)
}

func t2s(nn []int) int {
	var pairs int
	for _, n := range nn {
		for _, m := range nn {
			if n == m {
				continue
			}
			if n+m == T {
				pairs++
			}
		}
	}

	return pairs / 2
}
