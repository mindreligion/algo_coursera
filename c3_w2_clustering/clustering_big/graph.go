package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type graph struct {
	vs []uint32
	b int
}
func readGraph(file string) graph {
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
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	vs := make([]uint32, n)
	i := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		line = strings.Trim(line, " ")
		parts := strings.Split(line, " ")
		if len(parts) != b {
			panic("wrong parts number " + line)
		}
		var v uint32
		for j := len(parts) - 1; j >= 0; j-- {
			if parts[j] == "1" {
				v = v | 1<<(len(parts) - 1 - j)
			}
		}
		vs[i] = v
		i++
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	if i != n {
		fmt.Println("expected", n, "got:", i)
		panic("vertices count not match expected number")
	}

	return graph{
		vs: vs,
		b: b,
	}
}