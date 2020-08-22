package main

import (
	"math"
)

func mst(g graph) int64 {
	//start from 0 vertex
	h := &heap{}
	// put all vertices in the heap with corresponding cost of adding them
	for i := 1; i < len(g.vs); i++ {
		min := int64(math.MaxInt64)
		for _, e := range g.vs[i].es {
			if g.es[e].v1 == 0 || g.es[e].v2 == 0 {
				// vertex connected to 0 vertex
				if g.es[e].cost < min {
					min = g.es[e].cost
				}
			}
		}
		node := heapNode{
			v:    i,
			cost: min,
		}
		h.Put(node, g)
	}

	var sum int64
	for v, ok := h.Pop(g); ok; v, ok = h.Pop(g) {
		sum += v.cost
		for _, e := range g.vs[v.v].es {
			to := g.es[e].v1
			if to == v.v {
				to = g.es[e].v2
			}
			if g.vs[to].heapIdx == nil {
				continue
			}
			removed, ok := h.Remove(*(g.vs[to].heapIdx), g)
			if !ok {
				panic("unable to remove node from heap, that should be there")
			}
			min := removed.cost
			if g.es[e].cost < min {
				min = g.es[e].cost
			}
			node := heapNode{
				v:    to,
				cost: min,
			}
			h.Put(node, g)
		}
	}

	return sum
}
