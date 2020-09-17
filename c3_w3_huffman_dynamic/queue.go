package main

type queue struct {
	first *qnode
	last *qnode
}

func newQueue() *queue{
	return &queue{}
}

type qnode struct {
	next *qnode
	tree *node
}

func (q *queue) next() *node {
	if q.first == nil {
		return nil
	}

	return q.first.tree
}

func (q *queue) consume() *node {
	if q.first == nil {
		return nil
	}

	f := q.first
	q.first = f.next
	return f.tree
}

func (q *queue) add(t *node) {
	qn := qnode {
		tree: t,
	}

	if q.first == nil {
		q.first = &qn
		q.last = &qn

		return
	}

	q.last.next = &qn
	q.last = &qn
}

func printQ(q *queue) {
	node := q.first
	for node != nil {
		node = node.next
	}
}