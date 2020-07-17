package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type graph []vertex

type vertex struct {
	out     []int
	in      []int
	scc     int
	visited bool
	done    bool
}

func readGraph(path string, g *graph) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	*g = make(graph, 0)
	for j, line := range lines {
		line = strings.Trim(line, "\r")
		line = strings.Trim(line, " ")
		if line == "" {
			continue
		}

		words := strings.Split(line, " ")
		if len(words) != 2 {
			log.Fatalf("line %v words number is not 2", j)
		}

		from, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatalf("line %v word %v, error %v", j, 0, err)
		}

		to, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatalf("line %v word %v, error %v", j, 1, err)
		}

		to--
		from--

		g.addInEdge(to, from)
		g.addOutEdge(from, to)
	}
}

func (g *graph) addOutEdge(i int, to int) {
	v := vertex{
		out: []int{to},
		scc: -1,
	}
	if len(*g) == i {
		*g = append(*g, v)
		return
	}

	if len(*g) > i {
		// vertex already exists
		(*g)[i].out = append((*g)[i].out, to)
		return
	}

	n := i - len(*g)
	for i := 0; i < n; i++ {
		*g = append(*g, vertex{scc: -1})
	}

	*g = append(*g, v)
}

func (g *graph) addInEdge(i int, from int) {
	v := vertex{
		in:  []int{from},
		scc: -1,
	}
	if len(*g) == i {
		*g = append(*g, v)
		return
	}

	if len(*g) > i {
		// vertex already exists
		(*g)[i].in = append((*g)[i].in, from)
		return
	}

	n := i - len(*g)
	for i := 0; i < n; i++ {
		*g = append(*g, vertex{scc: -1})
	}

	*g = append(*g, v)
}
