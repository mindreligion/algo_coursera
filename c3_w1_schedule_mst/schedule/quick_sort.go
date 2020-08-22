package main

func quickSort(in []job, pivotIdx func(in []job, less func(i, j job) bool) int, less func(i, j job) bool) []job {
	// BC
	if len(in) <= 1 {
		return in
	}

	pi := pivotIdx(in, less)
	p := in[pi]

	if pi != 0 {
		in[pi], in[0] = in[0], in[pi]
	}

	b := 1

	for i := 1; i < len(in); i++ {
		if less(in[i], p) {
			in[i], in[b] = in[b], in[i]
			b++
		}
	}

	pi = b - 1
	in[0], in[pi] = in[pi], in[0]
	if pi > 0 {
		quickSort(in[:pi], pivotIdx, less)
	}
	if pi < len(in)-1 {
		quickSort(in[pi+1:], pivotIdx, less)
	}

	return in
}

func median(in []job, less func(i, j job) bool) int {
	l := len(in) - 1
	if len(in) < 3 {
		return 0
	}

	mi := len(in) / 2
	if len(in)%2 == 0 {
		mi--
	}

	if less(in[0], in[mi]) {
		if less(in[mi], in[l]) {
			return mi
		}
		if less(in[0], in[l]) {
			return l
		}

		return 0
	}
	if less(in[0], in[l]) {
		return 0
	}
	if less(in[mi], in[l]) {
		return l
	}

	return mi
}