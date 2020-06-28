package main

type heap []heapNode

type heapNode struct {
	v       int
	PathLen int
}

func (h *heap) Put(v heapNode, g graph) int {
	*h = append(*h, v)
	return h.bubbleUp(len(*h)-1, g)
}

func (h *heap) Min() (heapNode, bool) {
	s := *h
	if len(s) == 0 {
		return heapNode{}, false
	}
	return s[0], true
}

func (h *heap) Pop(g graph) (heapNode, bool) {
	min, ok := h.Min()
	if !ok {
		return heapNode{}, false
	}

	h.placeLastToPosAndCut(0, g)
	h.bubbleDown(0, g)

	return min, true
}

func (h *heap) Remove(i int, g graph) bool {
	if i >= len(*h) {
		return false
	}
	if i == len(*h)-1 {
		*h = (*h)[:len(*h)-1]
		return true
	}

	h.placeLastToPosAndCut(i, g)
	i = h.bubbleUp(i, g)
	h.bubbleDown(i, g)
	return true
}

func (h *heap) bubbleDown(i int, g graph) int {
	for {
		li := leftChildIdx(i)
		if li >= len(*h) {
			return i
		}

		ri := rightChildIdx(i)
		if ri >= len(*h) {
			if h.less(li, i) {
				h.swap(li, i, g)
				i = li
			}
			return i
		}

		min := li
		if h.less(ri, min) {
			min = ri
		}

		if !h.less(min, i) {
			return i
		}

		h.swap(min, i, g)
		i = min
	}
}

func (h *heap) bubbleUp(i int, g graph) int {
	for pi, ok := parentIdx(i); ok; pi, ok = parentIdx(i) {
		if !h.less(i, pi) {
			return i
		}

		h.swap(pi, i, g)
		i = pi
	}

	return i
}

func leftChildIdx(i int) int {
	return 2*i + 1
}

func rightChildIdx(i int) int {
	return 2*i + 2
}

func parentIdx(i int) (int, bool) {
	if i <= 0 {
		return 0, false
	}

	return (i - 1) / 2, true
}

func (h *heap) swap(i, j int, g graph) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	g[(*h)[i].v].heapIdx, g[(*h)[j].v].heapIdx = i, j
}

func (h *heap) less(i, j int) bool {
	return (*h)[i].PathLen < (*h)[j].PathLen
}

func (h *heap) placeLastToPosAndCut(i int, g graph) {
	(*h)[i] = (*h)[len(*h)-1]
	g[(*h)[i].v].heapIdx = i
	*h = (*h)[:len(*h)-1]
}
