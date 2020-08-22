package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	min = -10000
	max = 10000
)

func main() {
	f, err := os.OpenFile("./2_sum.txt", os.O_RDONLY, 0)
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
	count := t2sBuiltIn(nn)
	t := time.Now().Sub(start)
	fmt.Println(count)
	fmt.Println()
	fmt.Println(t)
}

func t2sBuiltIn(nn []int) int {
	var count int
	ht := make(map[int]bool)
	for _, n := range nn {
		ht[n] = true
	}

tLoop:
	for t := min; t <= max; t++ {
		start := time.Now()
		for _, n := range nn {
			m := t - n
			if n == m {
				continue
			}
			if ht[m] {
				count++
				continue tLoop
			}
		}
		duration := time.Now().Sub(start)
		fmt.Printf("processed t=%v in %v\n", t, duration)
	}

	return count
}

//func t2s(nn []int) int {
//	var pairs int
//	ht := ht2.NewHT()
//	for i, n := range nn {
//		ht.Add(n)
//		fmt.Printf("proceeded %v\n", i)
//	}
//	for i, n := range nn {
//		m := T - n
//		if n == m {
//			continue
//		}
//		pairs += ht.Find(m)
//		fmt.Printf("proceeded %v\n", i)
//	}
//
//	return pairs / 2
//}
