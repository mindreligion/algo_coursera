package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) knapsack {
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

	parts := strings.Split(s.Text(), " ")
	if len(parts) != 2 {
		panic("expected 2 word in first line")
	}
	size, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	i := 0
	items := make([]item, n)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("expected 2 word in %v line", i))
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		w, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		items[i] = item{
			v: v,
			w: w,
		}
		i++
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	if i != n {
		fmt.Println("expected", n, "got:", i)
		panic("items number not match expected number")
	}

	return knapsack{
		s: size,
		items: items,
	}
}