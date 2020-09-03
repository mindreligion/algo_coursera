package main

import (
	"fmt"
	"testing"
)

func TestUF(t *testing.T) {
	// test union find
	uf := newUF(9)

	uf.union(0, 1)
	fmt.Println("union(0, 1)")
	fmt.Println(uf)

	uf.union(0, 2)
	fmt.Println("union(0, 2)")
	fmt.Println(uf)

	uf.union(0, 6)
	fmt.Println("union(0, 6)")
	fmt.Println(uf)

	uf.union(0, 6)
	fmt.Println("union(0, 6)")
	fmt.Println(uf)

	uf.union(3, 4)
	uf.union(3, 5)
	fmt.Println("union(3, 4)")
	fmt.Println("union(3, 5)")
	fmt.Println(uf)

	uf.union(7, 8)
	fmt.Println("union(7, 8)")
	fmt.Println(uf)

	uf.union(0, 7)
	fmt.Println("union(0, 7)")
	fmt.Println(uf)

	fmt.Println(uf.find(2))
	fmt.Println(uf)
	uf.union(5, 2)
	fmt.Println("union(5, 2)")
	fmt.Println(uf)

	uf.find(6)
	uf.find(0)
	uf.find(3)
	uf.find(5)
	fmt.Println(uf)
}
