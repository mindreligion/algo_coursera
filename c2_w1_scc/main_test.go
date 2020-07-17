package main

import (
	"testing"
)

func BenchmarkTop5Scc(b *testing.B) {
	g := make(graph, 0)
	readGraph("./scc.txt", &g)
	cg := make(graph, len(g))
	copy(cg, g)
	for i := 0; i < b.N; i++ {
		top5Scc(g)
	}
}