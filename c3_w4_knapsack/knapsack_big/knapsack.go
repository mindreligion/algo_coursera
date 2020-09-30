package main

import (
	"fmt"
	"strings"
)

type knapsack struct {
	s     int
	items []item
}

type item struct {
	v int
	w int
}

func (k *knapsack) solve() int {
	// init a
	maxX := k.s
	maxI := len(k.items)
	a := make(map[string]int)
	r := k.solveNX(maxI, maxX, a)
	fmt.Println("saved ", c)
	fmt.Println("a len ", len(a))
	return r
}

var c = 0
func (k *knapsack) solveNX(i, x int, a map[string]int) int {
	if v, ok := a[key(i,x)]; ok{
		return v
		c++
	}
	if i == 0 || x == 0 {
		a[key(i,x)] = 0
		return 0
	}
	curr :=k.items[i-1]
	excluding := k.solveNX(i-1, x, a)
	if x < curr.w {
		a[key(i,x)] = excluding
		return excluding
	}

	including := curr.v + k.solveNX(i-1, x-curr.w, a)

	a[key(i,x)] = including
	if a[key(i,x)] < excluding {
		a[key(i,x)] = excluding
	}

	return a[key(i,x)]
}

func printA(a [][]int) {
	if !printable {
		return
	}
	h := len(a[0])
	out := ""
	for j := h - 1; j >= 0; j-- {
		out += formatS(fmt.Sprintf("%v|", j))
		row := make([]string, len(a))
		for i := 0; i < len(a); i++ {
			row[i] = formatI(a[i][j])
		}
		out += strings.Join(row, "")
		out += "\n"
	}
	for i := 0; i < len(a)+1; i++ {
		out += formatS("----------")
	}
	out += "\n"
	out += formatS("")
	for i := 0; i < len(a); i++ {
		out += formatI(i)
	}
	fmt.Println(out)
	fmt.Println()
}

func formatI(i int) string {
	s := fmt.Sprintf("%v", i)
	sn := 10 - len(s)
	for x := 0; x < sn; x++ {
		s += " "
	}

	return s
}

func formatS(s string) string {
	sn := 10 - len(s)
	for x := 0; x < sn; x++ {
		s += " "
	}

	return s
}

func key(i, x int) string {
	return fmt.Sprintf("%v.%v", i, x)
}
