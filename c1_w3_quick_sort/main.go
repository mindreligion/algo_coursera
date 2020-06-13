package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var comparisons int64 = 0

func main() {
	b, err := ioutil.ReadFile("./quick_sort_input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	input := make([]int, 0)
	for _, line := range lines {
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		input = append(input, num)
	}
	inputA := make([]int, len(input))
	copy(inputA, input)
	quickSort(inputA, func(in []int) int { return 0 })
	fmt.Println("first element pivot: ", comparisons)
	comparisons = 0

	inputB := make([]int, len(input))
	copy(inputB, input)
	quickSort(inputB, func(in []int) int { return len(in) - 1 })
	fmt.Println("last element pivot: ", comparisons)
	comparisons = 0

	quickSort(input, median)
	fmt.Println("median element pivot: ", comparisons)
}

func quickSort(in []int, pivotIdx func(in []int) int) []int {
	// BC
	if len(in) <= 1 {
		return in
	}

	comparisons += int64(len(in)) - 1

	pi := pivotIdx(in)
	p := in[pi]

	if pi != 0 {
		in[pi], in[0] = in[0], in[pi]
	}

	b := 1

	for i := 1; i < len(in); i++ {
		if in[i] < p {
			in[i], in[b] = in[b], in[i]
			b++
		}
	}

	pi = b - 1
	in[0], in[pi] = in[pi], in[0]
	if pi > 0 {
		quickSort(in[:pi], pivotIdx)
	}
	if pi < len(in)-1 {
		quickSort(in[pi+1:], pivotIdx)
	}

	return in
}

func median(in []int) int {
	l := len(in) - 1
	if len(in) < 3 {
		return 0
	}
	mi := len(in) / 2
	if len(in)%2 == 0 {
		mi--
	}
	if in[0] < in[mi] {
		if in[mi] < in[l] {
			return mi
		}
		if in[0] < in[l] {
			return l
		}

		return 0
	}
	if in[0] < in[l] {
		return 0
	}
	if in[mi] < in[l] {
		return l
	}

	return mi

}
