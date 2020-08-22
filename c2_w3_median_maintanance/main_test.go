package main

import "testing"

func TestHeap(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int
		maxRes   []int
		minRes   []int
	}{
		{
			sequence: []int{7, 9, 15, 1, 4},
			maxRes:   []int{15, 9, 7, 4, 1},
			minRes:   []int{1, 4, 7, 9, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minH := minHeap{}
			maxH := maxHeap{}
			for _, i := range tt.sequence {
				minH.Put(i)
				maxH.Put(i)
			}
			if !sliceEqual(heapToSlice(&minH), tt.minRes) {
				t.Error("min heap is failed")
			}

			if !sliceEqual(heapToSlice(&maxH), tt.maxRes) {
				t.Error("max heap is failed")
			}
		})
	}
}

func sliceEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
