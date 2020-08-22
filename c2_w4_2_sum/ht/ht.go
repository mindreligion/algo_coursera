package ht

import (
	"fmt"
	"github.com/mindreligion/algo_coursera/c2_w4_2_sum/prime"
)

const size = 1e7

type HT struct {
	size int
	data []*ll
}

func NewHT() *HT {
	n := prime.PrimeN(size)
	data := make([]*ll, n)
	for i := range data {
		data[i] = NewLL()
		fmt.Println("processed ", i)
	}
	return &HT{
		size: 0,
		data: data,
	}
}

func (h *HT) hash(i int) int {
	return i % len(h.data)
}

func (h *HT) Add(i int) {
	idx := h.hash(i)
	fmt.Println(i, idx)
	if ok, node := h.data[idx].Find(i); ok {
		node.Increment()
		return
	}
	h.data[idx].Add(i)
}

func (h *HT) Find(i int) int {
	if ok, node := h.data[h.hash(i)].Find(i); ok {
		return node.Val()
	}

	return 0
}
