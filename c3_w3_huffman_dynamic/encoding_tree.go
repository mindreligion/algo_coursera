package main

type node struct {
	index *int
	weight int
	left *node
	right *node
}

func traverse(root *node) []string {
	rm := traverseRec(root, "")
	enc :=make([]string, len(rm))
	for i, w := range rm {
		enc[i] = w
	}

	return enc
}

func traverseRec(root *node, prefix string) map[int]string {
	if root.left == nil {
		if root.index == nil {
			panic("index == nil and left == nil")
		}
		return map[int]string{*root.index: prefix}
	}

	if root.right == nil {
		panic("right branch == nil shouldn't happen")
	}

	lr := traverseRec(root.left, prefix + "0")
	rr := traverseRec(root.right, prefix + "1")

	for i, w := range rr {
		lr[i] = w
	}
	return lr
}