package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type graph []vertex

type vertex struct {
	i       int
	es      []edge
	pathLen int
	heapIdx int
}

type edge struct {
	head int
	len  int
}

func readGraph(path string) graph {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	g := make(graph, 0)
	for j, line := range lines {
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}

		words := strings.Split(line, "\t")
		if len(words) == 0 {
			panic("empty words list")
		}
		vid, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatalf("line %v word %v, error %v", j, 0, err)
		}

		v := vertex{
			i:       vid - 1,
			pathLen: -1,
			heapIdx: -1,
		}

		// parse edges
		for i := 1; i < len(words); i++ {
			if words[i] == "" {
				continue
			}
			parts := strings.Split(words[i], ",")
			if len(parts) != 2 {
				log.Fatalf("line %v word %v: %v, error %v", j, i, words[i], err)
			}
			head, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("line %v word %v, part one %v, error %v", j, i, parts[0], err)
			}
			head--

			length, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("line %v word %v, part two %v, error %v", j, i, parts[1], err)
			}

			e := edge{
				len:  length,
				head: head,
			}

			v.es = append(v.es, e)
		}

		g.putInPos(v, v.i)
	}

	return g
}

func (g *graph) putInPos(v vertex, i int) {
	if len(*g) == i {
		*g = append(*g, v)
		return
	}

	if len(*g) > i {
		log.Printf("WARNING: put vertex to graph len %v in position %v", len(*g), i)
		(*g)[i] = v
		return
	}

	log.Printf("WARNIGN: put vertex to graph len %v in position %v", len(*g), i)
	n := i - len(*g)
	for i := 0; i < n; i++ {
		*g = append(*g, vertex{})
	}

	*g = append(*g, v)
}
