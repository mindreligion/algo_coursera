package main

import (
	"flag"
	"fmt"
)

var f *string
var k *int

func init() {
	f = flag.String("f", "main", "data file")
	k = flag.Int("k", 4, "number of clusters")
	flag.Parse()
}

func main() {
	var file string
	switch *f {
	case "main":
		file = "./clustering.txt"
	case "test":
		file = "./clustering_test.txt"
	case "test2":
		file = "./clustering_test1.txt"

	default:
		panic("unknown file" + *f)
	}

	g := readGraph(file)
	fmt.Println(g.maxSpacingKClustering(*k))
}
