package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type vertexCfg struct {
	id       int
	adjacent []int
}

type graph struct {
	vs []*vertex
	es []*edge
}

type vertex struct {
	i  int
	es []*edge
}

type edge struct {
	i  int
	v1 *vertex
	v2 *vertex
}

func main() {
	b, err := ioutil.ReadFile("./min_cut_input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	var gCfg []vertexCfg
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
		v := vertexCfg{
			id: vid,
		}

		for i := 1; i < len(words); i++ {
			if words[i] == "" {
				continue
			}

			num, err := strconv.Atoi(words[i])
			if err != nil {
				log.Fatalf("line %v word %v, error %v", j, i, err)
			}

			v.adjacent = append(v.adjacent, num)
		}

		gCfg = append(gCfg, v)
	}

	minCut := minCut(gCfg, len(gCfg))
	fmt.Println(minCut)
}

func buildGraph(cfg []vertexCfg) graph {
	g := graph{
		vs: make([]*vertex, len(cfg)),
	}
	for _, v := range cfg {
		i := v.id - 1
		g.vs[i] = &vertex{i: i}
	}
	for _, v := range cfg {
		i := v.id - 1
		for _, e := range v.adjacent {
			j := e - 1
			if j >= i {
				eid := len(g.es)
				edge := edge{
					i:  eid,
					v1: g.vs[i],
					v2: g.vs[j],
				}
				g.es = append(g.es, &edge)
				g.vs[i].es = append(g.vs[i].es, &edge)
				if j > i {
					g.vs[j].es = append(g.vs[j].es, &edge)
				}
			}
		}
	}

	return g
}

func printG(g graph) {
	fmt.Println("Vertices:")
	for _, v := range g.vs {
		fmt.Println("\tVertex: ", v.i)
		fmt.Println("\tEdges to:")
		for _, e := range v.es {
			to := e.v1.i
			if to == v.i {
				to = e.v2.i
			}
			fmt.Println("\t\t", to)
		}
		fmt.Println()
	}
	fmt.Println("Edges:")
	for _, e := range g.es {
		fmt.Println("\tEdge: ", e.i)
		fmt.Println("\tConnects vertices:")
		fmt.Println("\t\t", e.v1.i, " ", e.v2.i)
		fmt.Println()
	}
}

func minCut(cfg []vertexCfg, n int) int {
	var min int
	for i := 0; i < n; i++ {
		cut := randContractionMC(buildGraph(cfg))
		if i == 0 {
			min = cut
			continue
		}
		if cut < min {
			min = cut
		}
	}
	return min
}

func randContractionMC(g graph) int {
	for len(g.vs) > 2 {
		// Step 1
		e := g.es[int(r.Int31n(int32(len(g.es))))]

		// Step 2
		v1 := e.v1
		v2 := e.v2
		for _, v2e := range v2.es {
			if v2e.v1 == v2 {
				v2e.v1 = v1
				continue
			}
			if v2e.v2 == v2 {
				v2e.v2 = v1
				continue
			}
			log.Fatal("vertex points to the edge, that not connects this vertex")
		}

		// Step 3
		v1.es = addUniqueList(v1.es, v2.es)
		// Step 3.1
		for i := range v2.es {
			v2.es[i] = nil
		}

		// Step 3.2
		last := len(g.vs) - 1
		if v2.i != last {
			g.vs[last].i = v2.i

			// Step 3.3
			g.vs[v2.i] = g.vs[last]
		}

		// Step 3.4
		g.vs[last] = nil

		// Step 3.5
		g.vs = g.vs[:last]

		// Step 4
		for i := 0; i < len(v1.es); i++ {
			e := v1.es[i]
			if e.v1 != e.v2 {
				continue
			}
			e.v1 = nil
			e.v2 = nil

			// remove from vertex's edge list
			last := len(v1.es) - 1
			needRedo := false
			if i != last {
				needRedo = true
				v1.es[i] = v1.es[last]
			}
			v1.es[last] = nil
			v1.es = v1.es[:last]

			// remove edge from edge slice
			last = len(g.es) - 1
			if e.i != last {
				g.es[last].i = e.i
				g.es[e.i] = g.es[last]
			}
			g.es[last] = nil
			g.es = g.es[:last]
			if needRedo {
				i--
			}
		}
	}

	return len(g.es)
}

func addUniqueList(es []*edge, add []*edge) []*edge {
	for i := range add {
		es = addUnique(es, add[i])
	}

	return es
}

func addUnique(es []*edge, e *edge) []*edge {
	for i := range es {
		if es[i] == e {
			return es
		}
	}
	return append(es, e)
}
