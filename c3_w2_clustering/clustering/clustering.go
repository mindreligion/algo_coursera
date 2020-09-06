package main

func (g *graph) maxSpacingKClustering(k int) int {
	g.sortEdges()
	uf := newUF(g.n)
	i := 0
	for uf.c > k {
		uf.union(g.es[i].v1, g.es[i].v2)
		i++
	}

	for ; i < len(g.es); i++ {
		l1 := uf.find(g.es[i].v1)
		l2 := uf.find(g.es[i].v2)
		if l1 != l2 {
			return g.es[i].cost
		}
	}

	return 0
}
