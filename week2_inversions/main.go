package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
var counter = 0

func main() {
	b, err := ioutil.ReadFile("./inversions_input.txt")
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
	fmt.Println(counter, inversionsCount(input))
}

func inversionsCount(a []int) int {
	n, _ := inversionsRec(a)
	return n
}

func inversionsRec(a []int) (int, []int) {
	counter++

	if len(a) == 1 {
		return 0, a
	}
	mid := len(a)/2
	invL, l := inversionsRec(a[:mid])
	invR, r := inversionsRec(a[mid:])
	return mergeInv(invL, invR, l, r)
}

func mergeInv(invL, invR int, l, r []int) (int, []int) {
	invCross := 0
	merged := make([]int, 0, len(l) + len(r))
	for i, j := 0, 0;;{
		if l[i] <= r[j] {
			merged = append(merged, l[i])
			i++
			if i == len(l) {
				merged = append(merged,r[j:]...)
				break
			}
			continue
		}

		merged = append(merged, r[j])
		j++
		invCross += len(l)-i
		if j == len(r) {
			merged = append(merged, l[i:]...)
			break
		}

	}
	return invCross + invL + invR, merged
}