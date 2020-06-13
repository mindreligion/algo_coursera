package main

import "testing"

func Test_median(t *testing.T) {

	tests := []struct {
		name string
		in   []int
		want int
	}{
		{
			name: "1",
			in:   []int{5,7,8,9},
			want: 1,
		},
		{
			name: "2",
			in:   []int{5,7,8,9, 10},
			want: 2,
		},
		{
			name: "1",
			in:   []int{8, 7, 5, 9, 10},
			want: 0,
		},
		{
			name: "1",
			in:   []int{10, 5, 7, 8},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median(tt.in); got != tt.want {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}
