package main

import "fmt"

func main() {
	g := readGraph("./min_cut_input_test2.txt")
	s := 0
	dijkstraMinPath(g, s)
	printMinPath(g, s)
}

func printMinPath(g graph, s int) {
	for _, v := range g {
		fmt.Printf("Min path from vertex %v to %v is %v\n", s, v.i, v.pathLen)
	}
}
