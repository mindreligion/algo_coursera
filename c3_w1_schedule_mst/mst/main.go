package main

import (
	"flag"
)

var f *string

func init() {
	f = flag.String("f", "main", "data file")
	flag.Parse()
}

func main() {
	var file string
	switch *f {
	case "main":
		file = "./edges.txt"
	case "test":
		file = "./edges_test.txt"
	case "test2":
		file = "./edges_test2.txt"

	default:
		panic("unknown file" + *f)
	}

	g := readGraph(file)
	println(mst(g))
}
