package main

import (
	"bufio"
	"fmt"
	ht2 "github.com/mindreligion/algo_coursera/c2_w4_2_sum/ht"
	"os"
	"strconv"
	"time"
)

const T = 666

func main() {
	f, err := os.OpenFile("./2_sum_test1.txt", os.O_RDONLY, 0)
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

func t2sBuiltIn(nn []int) int {
	var pairs int
	ht := make(map[int]int)
	for _, n := range nn {
		ht[n]++
	}
	for _, n := range nn {
		m := T - n
		if n == m {
			continue
		}
		if ht[m] > 0 {
			pairs += ht[m]
		}
	}

	return pairs / 2
}

func t2s(nn []int) int {
	var pairs int
	ht := ht2.NewHT()
	for i, n := range nn {
		ht.Add(n)
		fmt.Printf("proceeded %v\n", i)
	}
	for i, n := range nn {
		m := T - n
		if n == m {
			continue
		}
		pairs += ht.Find(m)
		fmt.Printf("proceeded %v\n", i)
	}

	return pairs / 2
}
