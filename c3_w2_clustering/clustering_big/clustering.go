package main

import "fmt"

type mutation struct {
	p int
	v uint32
}

func (g *graph) maxK3SpacingClustering() int {
	mm := make([]mutation, (g.b + 1) * len(g.vs))
	for i := range g.vs {
		s := i * (g.b + 1)
		mm[s] = mutation{
			p: i,
			v: g.vs[i],
		}
		for j := 0; j < g.b; j++ {
			mm[s+j+1] = mutation{
				p: i,
				v: g.vs[i] ^ 1<<j,
			}
		}
	}

	quickSort(mm)
	uf := newUF(len(g.vs))
	for i := 1; i < len(mm); i++ {
		if mm[i].v == mm[i-1].v {
			if mm[i].p != mm[i-1].p {
				uf.union(mm[i].p, mm[i-1].p)
			} else {
				fmt.Println(mm[i], mm[i-1])
				panic("weird stuff happen")
			}
		}
	}

	return uf.c
}
