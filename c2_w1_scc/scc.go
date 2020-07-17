package main

func (g graph) RDfsLoop() []int {
	var order []int
	for i, v := range g {
		if v.visited {
			continue
		}
		g.RDfs(i, &order)
	}
	return order
}

func (g graph) RDfs(v int, order *[]int) {
	if g[v].visited {
		return
	}
	g[v].visited = true
	for _, i := range g[v].in {
		g.RDfs(i, order)
	}
	*order = append(*order, v)
}

func (g graph) DfsLoop(order []int) {
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		g.Dfs(v, v)
	}
}

func (g graph) Dfs(v int, mark int) {
	if g[v].scc >= 0 {
		return
	}
	g[v].scc = mark
	for _, i := range g[v].out {
		g.Dfs(i, mark)
	}
}

func (g graph) RDfsLoopStack() []int {
	var order []int
	for i, v := range g {
		if v.visited {
			continue
		}
		g.RDfs(i, &order)
	}
	return order
}

func (g graph) RDfsStack(v int, order *[]int) {
	if g[v].visited {
		return
	}
	var stack []int
	stack = append(stack, v)
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		if g[v].done {
			continue
		}

		if g[v].visited {
			*order = append(*order, v)
			g[v].done = true
		}

		for _, i := range g[v].in {
			if !g[i].visited {
				stack = append(stack, v)
			}
		}
		g[v].visited = true
	}

}
