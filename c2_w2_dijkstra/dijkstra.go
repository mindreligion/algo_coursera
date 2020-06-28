package main

import "log"

func dijkstraMinPath(g graph, s int) {
	if len(g) <= s {
		log.Fatalf("graf of len %v doesn't contain vertex %v to use as source", len(g), s)
	}

	h := heap{heapNode{
		v:       s,
		PathLen: 0,
	}}
	for node, ok := h.Pop(g); ok; node, ok = h.Pop(g) {
		s := node.v
		g[s].pathLen = node.PathLen
		g[s].heapIdx = -1

		// for each edge with tail in s, and path to head not yet calculated put edge head vertex to the heap with corresponding path len
		// (if it's already here put/left shortest path)
		for _, e := range g[s].es {
			if g[e.head].pathLen >= 0 {
				// skip edge pointing to the vertex for which path is calculated already.
				continue
			}

			l := g[s].pathLen + e.len
			hi := g[e.head].heapIdx
			if hi >= 0 {
				// head vertex already in heap
				if h[hi].v != e.head {
					log.Fatal("vertex points on heap node that points on different vertex")
				}
				if h[hi].PathLen <= l {
					// path to e.head that already in heap is shorter or equal then path from s, skip this edge and continue.
					continue
				}
				h.Remove(hi, g)
			}

			g[e.head].heapIdx = h.Put(heapNode{
				v:       e.head,
				PathLen: l,
			}, g)
		}
	}
}
