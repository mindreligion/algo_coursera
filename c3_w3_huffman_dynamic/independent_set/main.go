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
	var nodesToCheck []int
	switch *f {
	case "main":
		file = "./mwis.txt"
		nodesToCheck = []int{1, 2, 3, 4, 17, 117, 517, 997}
	case "test":
		file = "./mwis_test.txt"
		nodesToCheck = []int{1, 2, 3,  4, 5, 6, 7, 8}

	default:
		panic("unknown file" + *f)
	}

	ww := readInput(file)
	sum, is := mwis(ww)
	fmt.Println(sum)
	fmt.Println(inIS(is, nodesToCheck))
}

func inIS(is []bool, nodes []int) string {
	res := ""
	for _, node := range nodes {
		i := node - 1
		if is[i] {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}
