package main

import "testing"

func Test_inversionsCount(t *testing.T) {
	tests := []struct {
		name string
		a []int
		want int
	}{
		{
			name: "one inv",
			a: []int{1, 2, 3, 5, 4},
			want: 1,
		},
		{
			name: "6 inv",
			a: []int{7, 2, 5, 6, 3},
			want: 6,
		},
		{
			name: "4 inv",
			a: []int{2, 1, 4, 5, 9, 3},
			want: 4,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inversionsCount(tt.a); got != tt.want {
				t.Errorf("inversionsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}