package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type graph struct {
	//vs []vertex
	es []edge
	n int
}

//type vertex struct {
//	es []int
//}

type edge struct {
	v1   int
	v2   int
	cost int
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
	if len(parts) != 1 {
		panic("expected 1 word in first line")
	}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	//vs := make([]vertex, n)
	m := (n-1)*n/2
	es := make([]edge, m)
	i := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			panic("wrong parts number " + line)
		}
		v1, err := strconv.Atoi(parts[0])
		v1--
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(parts[1])
		v2--
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}

		es[i] = edge{
			v1:   v1,
			v2:   v2,
			cost: c,
		}
		//vs[v1].es = append(vs[v1].es, i)
		//vs[v2].es = append(vs[v2].es, i)
		i++
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	if i != m {
		fmt.Println("expected", m, "got:", i)
		panic("edges count not match expected number")
	}

	return graph{
		//vs: vs,
		es: es,
		n: n,
	}
}

func (g *graph) sortEdges() {
	quickSort(g.es)
}