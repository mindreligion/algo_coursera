package ht

import "fmt"

const stopLoop = 1e6

type ll struct {
	Head *node
}

type node struct {
	next *node
	key  int
	val  int
}

func NewLL() *ll {
	return &ll{}
}

func (l *ll) Add(i int) {
	l.Head = &node{
		next: l.Head,
		key:  i,
		val:  1,
	}
}

func (l *ll) Find(i int) (bool, *node) {
	ll := l.Head
	c := 0
	for ll != nil {
		c++
		fmt.Println(c)
		if c > stopLoop {
			panic("stop fucking loop" + fmt.Sprintf("%v, %v", ll, ll.next))
		}
		if ll.key == i {
			return true, ll
		}
		ll = ll.next
	}

	return false, nil
}

func (n *node) Increment() {
	n.val++
}

func (n *node) Val() int {
	return n.val
}
