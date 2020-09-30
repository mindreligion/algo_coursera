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
	maxX := k.s
	maxI := len(k.items)

	a := make([][]int, maxI+1)
	for i := 0; i <= maxI; i++ {
		a[i] = make([]int, maxX+1)
		for x := 0; x <= maxX; x++ {
			if x == 0 || i == 0 {
				continue
			}
			excluding := a[i-1][x]
			if x < k.items[i-1].w {
				a[i][x] = excluding
				continue
			}
			including := k.items[i-1].v + a[i-1][x-k.items[i-1].w]

			a[i][x] = including
			if a[i][x] < excluding {
				a[i][x] = excluding
			}
		}
	}

	printA(a)

	return a[maxI][maxX]
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
