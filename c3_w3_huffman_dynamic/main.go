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
		file = "./huffman.txt"
	case "test":
		file = "./huffman_test.txt"

	default:
		panic("unknown file" + *f)
	}

	ww := readInput(file)
	enc := huffmanEncoding(ww)
	fmt.Println(minMax(enc))
}
