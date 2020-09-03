package main

import "fmt"

func newUF(n int) uf {
	l := make([]int, n)
	for i := range l {
		l[i] = i
	}

	return uf{
		l: l,
		r: make([]int, n),
		c: n,
	}
}

type uf struct {
	l []int
	r []int
	c int
}

func (uf *uf) union(a int, b int) {
	la := uf.find(a)
	lb := uf.find(b)
	if la == lb {
		return
	}
	uf.c--
	if uf.r[la] > uf.r[lb] {
		uf.l[lb] = la
		return
	}

	uf.l[la] = lb
	if uf.r[lb] == uf.r[la] {
		uf.r[lb]++
	}
}

func (uf *uf) find(i int) int {
	if i < 0 || i > len(uf.l) {
		panic("vertex index out of range " + fmt.Sprint(i))
	}
	v := i
	var l int
	for {
		l = uf.l[v]
		if v == l {
			break
		}
		v = l
	}
	sl := l
	// path compression
	v = i
	for {
		l := uf.l[v]
		if l == sl {
			break
		}
		uf.l[v] = sl
		v = l
	}
	return sl
}
