package main

import (
	"flag"
	"fmt"
)

func init() {
	f = flag.String("f", "main", "data file")
	flag.Parse()
}

var (
	f         *string
	printable bool
)

func main() {
	var file string
	switch *f {
	case "main":
		file = "./knapsack1.txt"
	case "test":
		file = "./knapsack_test.txt"
		printable = true
	case "test2":
		file = "./knapsack_test2.txt"
		printable = true
	case "big":
		file = "./knapsack_big.txt"

	default:
		panic("unknown file" + *f)
	}

	ks := readInput(file)
	fmt.Println(ks.solve())
}
