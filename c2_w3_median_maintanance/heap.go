package main

type heapI interface {
	Put(v int) int
	Root() (int, bool)
	Pop() (int, bool)
}

type minHeap []int

func (h *minHeap) Put(v int) int {
	*h = append(*h, v)
	return h.bubbleUp(len(*h) - 1)
}

func (h *minHeap) Root() (int, bool) {
	s := *h
	if len(s) == 0 {
		return 0, false
	}
	return s[0], true
}

func (h *minHeap) Pop() (int, bool) {
	min, ok := h.Root()
	if !ok {
		return 0, false
	}

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.bubbleDown(0)

	return min, true
}

func (h *minHeap) bubbleDown(i int) int {
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

func (h *minHeap) bubbleUp(i int) int {
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

type maxHeap []int

func (h *maxHeap) Put(v int) int {
	*h = append(*h, v)
	return h.bubbleUp(len(*h) - 1)
}

func (h *maxHeap) Root() (int, bool) {
	s := *h
	if len(s) == 0 {
		return 0, false
	}
	return s[0], true
}

func (h *maxHeap) Pop() (int, bool) {
	min, ok := h.Root()
	if !ok {
		return 0, false
	}

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.bubbleDown(0)

	return min, true
}

func (h *maxHeap) bubbleDown(i int) int {
	for {
		li := leftChildIdx(i)
		if li >= len(*h) {
			return i
		}

		ri := rightChildIdx(i)
		if ri >= len(*h) {
			if (*h)[li] > (*h)[i] {
				(*h)[li], (*h)[i] = (*h)[i], (*h)[li]
				i = li
			}
			return i
		}

		max := li
		if (*h)[max] < (*h)[ri] {
			max = ri
		}

		if (*h)[max] <= (*h)[i] {
			return i
		}

		(*h)[max], (*h)[i] = (*h)[i], (*h)[max]
		i = max
	}
}

func (h *maxHeap) bubbleUp(i int) int {
	for pi, ok := parentIdx(i); ok; pi, ok = parentIdx(i) {
		if (*h)[pi] >= (*h)[i] {
			return i
		}

		(*h)[pi], (*h)[i] = (*h)[i], (*h)[pi]
		i = pi
	}

	return i
}
