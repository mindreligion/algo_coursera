package main

func quickSort(in []char) []char {
	// BC
	if len(in) <= 1 {
		return in
	}

	pi := median(in)
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
		quickSort(in[:pi])
	}
	if pi < len(in)-1 {
		quickSort(in[pi+1:])
	}

	return in
}

func median(in []char) int {
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

func less(i, j char) bool {
	return i.w < j.w
}

