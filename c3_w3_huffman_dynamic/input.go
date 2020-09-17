package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(file string) []int {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	if !s.Scan() {
		if s.Err() != nil {
			panic(err)
		}
		panic("empty file")
	}

	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}

	i := 0
	ww := make([]int, n)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		w, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ww[i] = w
		i++
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	if i != n {
		fmt.Println("expected", n, "got:", i)
		panic("edges count not match expected number")
	}

	return ww
}