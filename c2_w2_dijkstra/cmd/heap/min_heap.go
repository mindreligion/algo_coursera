package main

type heap []int

func (h *heap) Put(v int) int {
	*h = append(*h, v)
	return h.bubbleUp(len(*h) - 1)
}

func (h *heap) Min() (int, bool) {
	s := *h
	if len(s) == 0 {
		return 0, false
	}
	return s[0], true
}

func (h *heap) Pop() (int, bool) {
	min, ok := h.Min()
	if !ok {
		return 0, false
	}

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.bubbleDown(0)

	return min, true
}

func (h *heap) Remove(i int) bool {
	if i >= len(*h) {
		return false
	}

	(*h)[i] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	i = h.bubbleUp(i)
	h.bubbleDown(i)
	return true
}

func (h *heap) bubbleDown(i int) int {
	for {
		li := leftChildIdx(i)
		if li >= len(*h) {
			return i
		}

		ri := rightChildIdx(i)
		if ri >= len(*h) {
			if (*h)[li] < (*h)[i] {
				(*h)[li], (*h)[i] = (*h)[i], (*h)[li]
				i = li
			}
			return i
		}

		min := li
		if (*h)[min] > (*h)[ri] {
			min = ri
		}

		if (*h)[min] >= (*h)[i] {
			return i
		}

		(*h)[min], (*h)[i] = (*h)[i], (*h)[min]
		i = min
	}
}

func (h *heap) bubbleUp(i int) int {
	for pi, ok := parentIdx(i); ok; pi, ok = parentIdx(i) {
		if (*h)[pi] <= (*h)[i] {
			return i
		}

		(*h)[pi], (*h)[i] = (*h)[i], (*h)[pi]
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
