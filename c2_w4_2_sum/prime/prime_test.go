package prime

import "testing"

func TestPrimeN(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "18->17",
			n:    18,
			want: 17,
		},
		{
			name: "20->19",
			n:    20,
			want: 19,
		},
		{
			name: "2->1",
			n:    2,
			want: 1,
		},
		{
			name: "81->79",
			n:    81,
			want: 79,
		},
		{
			name: "1000->997",
			n:    1000,
			want: 997,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeN(tt.n); got != tt.want {
				t.Errorf("PrimeN() = %v, want %v", got, tt.want)
			}
		})
	}
}
