package main

import (
	"flag"
	"fmt"
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
		file = "./clustering_big.txt"
	case "test":
		file = "./clustering_big_test.txt"
	default:
		panic("unknown file" + *f)
	}

	g := readGraph(file)
	fmt.Println(g.maxK3SpacingClustering())
}
