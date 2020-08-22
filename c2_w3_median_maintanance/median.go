package main

type median struct {
	minh *minHeap
	maxh *maxHeap
}

func (m *median) Put(i int) {
	max, _ := m.maxh.Root()
	if i > max {
		m.minh.Put(i)
	} else {
		m.maxh.Put(i)
	}

	// balance heaps
	if len(*m.maxh)-len(*m.minh) > 1 {
		i, _ := m.maxh.Pop()
		m.minh.Put(i)
		return
	}

	if len(*m.minh) > len(*m.maxh) {
		i, _ := m.minh.Pop()
		m.maxh.Put(i)
	}
}

func (m *median) Median() int {
	i, _ := m.maxh.Root()
	return i
}
