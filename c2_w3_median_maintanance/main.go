package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.OpenFile("./median.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	m := median{
		minh: &minHeap{},
		maxh: &maxHeap{},
	}
	msum := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		err := s.Err()
		if err != nil {
			panic(err)
		}
		line := s.Text()
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(line)
			panic(err)
		}
		m.Put(i)
		msum += m.Median()
	}

	fmt.Println(msum % 10000)
}

func printHeap(heap heapI) {
	for _, v := range heapToSlice(heap) {
		fmt.Println(v)
	}
}

func heapToSlice(heap heapI) []int {
	var s []int
	for i, ok := heap.Pop(); ok; i, ok = heap.Pop() {
		s = append(s, i)
	}
	return s
}
