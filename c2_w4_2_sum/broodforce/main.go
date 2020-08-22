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
	sums := t2s(nn)
	t := time.Now().Sub(start)
	fmt.Println(sums)
	fmt.Println()
	fmt.Println(t)
}

func t2s(nn []int) int {
	var count int
tLoop:
	for t := min; t <= max; t++ {
		start := time.Now()

		for _, n := range nn {
			for _, m := range nn {
				if n == m {
					continue
				}
				if n+m == t {
					count++
					continue tLoop
				}
			}
		}
		duration := time.Now().Sub(start)
		fmt.Printf("processed t=%v in %v\n", t, duration)
	}

	return count
}
