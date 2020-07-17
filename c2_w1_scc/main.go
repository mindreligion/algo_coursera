package main

import (
	"fmt"
	"sort"
)

func main() {
	go NewMonitor(50)
	g := make(graph, 0)
	readGraph("./scc_test.txt", &g)
	//fmt.Println(g)
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
	fmt.Println(sccList[:l])
}
