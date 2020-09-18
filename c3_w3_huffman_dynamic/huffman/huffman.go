package main

import "math"

func huffmanEncoding(ww []int) []string {
	cc := make([]char, len(ww))
	for i, w := range ww {
		cc[i] = char{
			index: i,
			w:     w,
		}
	}

	quickSort(cc)
	merged := newQueue()
	single := newQueue()
	for i := range cc {
		single.add(&node{
			index:  &cc[i].index,
			weight: cc[i].w,
		})
	}
	t := buildTree(single, merged)
	return traverse(t)
}

func buildTree(single *queue, merged *queue) *node {
	for {
		t1 := min(single, merged)
		t2 := min(single, merged)
		if t2 == nil {
			return t1
		}
		merged.add(&node{
			weight: t1.weight + t2.weight,
			left:   t1,
			right:  t2,
		})
	}
}

func min(single *queue, merged *queue) *node {
	s := single.next()
	m := merged.next()
	min := s
	if s != nil && (m == nil || s.weight < m.weight) {
		single.consume()
	} else {
		min = m
		merged.consume()
	}
	return min
}

type char struct {
	index int
	w     int
}

func minMax(enc []string) (int, int) {
	max := 0
	min := math.MaxInt32
	for i := range enc {
		if min > len(enc[i]) {
			min = len(enc[i])
		}
		if max < len(enc[i]) {
			max = len(enc[i])
		}
	}

	return min, max
}
