package main

import "fmt"

func mwis(p []int) (int, []bool) {
	if len(p) == 0 {
		return 0, nil
	}

	a := make([]int, len(p)+1)
	a[0] = 0
	a[1] = p[0]
	for i := 1; i < len(p); i++ {
		a[i+1] = a[i]
		a0 := a[i-1] + p[i]
		if a0 > a[i+1] {
			a[i+1] = a0
		}
	}
	sum := a[len(a)-1]

	//reverse
	is := make([]bool, len(a)-1)
	for i := len(a) - 1; i > 0; {
		if a[i] == a[i-1] {
			i--
			continue
		}

		if i == 1 || a[i] == a[i-2] + p[i-1] {
			is[i-1] = true
			i -= 2
			continue
		}

		panic(fmt.Sprintf("both cases missing %d", i))
	}

	return sum, is
}
