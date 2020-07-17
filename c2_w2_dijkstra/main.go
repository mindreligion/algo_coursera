package main

import (
	"fmt"
	"log"
)

func main() {
	g := readGraph("./dijkstra.txt")
	s := 0
	dijkstraMinPath(g, s)
	printMinPath(g, s)
}

func printMinPath(g graph, s int) {
	for i, v := range g {
		if i != v.i {
			log.Fatalf("index in array do not match vertex index i: %v, v.i %v\n", i, v.i)
		}

		//fmt.Printf("Min path from vertex %v to %v is %v\n", s+1, v.i+1, v.pathLen)
	}

	significant := []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	out := fmt.Sprintf("%d", g[significant[0]-1].pathLen)
	for i := 1; i < len(significant); i++ {
		out = fmt.Sprintf("%s,%d", out, g[significant[i]-1].pathLen)
	}
	fmt.Println(out)
}
