package main

import (
	"fmt"
	"sort"
)

func main() {
	g := make(graph, 0)
	readGraph("./scc_test2.txt", &g)
	//fmt.Println(g)
	fmt.Println(top5Scc(g))
}

func top5Scc(g graph) []int {
	//go NewMonitor(50)

	order := g.RDfsLoop()
	g.DfsLoop(order)
	sccs := make(map[int]int)
	for _, v := range g {
		sccs[v.scc]++
	}
	//fmt.Println(sccs)
	sccList := make([]int, len(sccs))
	i := 0
	for _, v := range sccs {
		sccList[i] = v
		i++
	}

	sort.Slice(sccList, func(i int, j int) bool { return sccList[i] > sccList[j] })
	l := 5
	if len(sccList) < 5 {
		l = len(sccList)
	}

	return sccList[:l]
}
