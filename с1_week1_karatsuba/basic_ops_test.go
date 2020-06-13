package main

import (
	"fmt"
	"testing"
)

func TestByteProduct(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "2*3=6",
			args: args{
				a: '2',
				b: '3',
			},
			want: "6",
		},
		{
			name: "7*8=56",
			args: args{
				a: '7',
				b: '8',
			},
			want: "56",
		},
		{
			name: "2*9=18",
			args: args{
				a: '2',
				b: '9',
			},
			want: "18",
		},
		{
			name: "9*2=18",
			args: args{
				a: '9',
				b: '2',
			},
			want: "18",
		},
		{
			name: "3*3=9",
			args: args{
				a: '3',
				b: '3',
			},
			want: "9",
		},
		{
			name: "0*7=0",
			args: args{
				a: '0',
				b: '7',
			},
			want: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteProduct(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("byteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSum(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0+0=0",
			args: args{
				a: "0",
				b: "0",
			},
			want: "0",
		},
		{
			name: "2332+0",
			args: args{
				a: "2332",
				b: "0",
			},
			want: "2332",
		},
		{
			name: "322323+8382892392",
			args: args{
				a: "322323",
				b: "8382892392",
			},
			want: fmt.Sprintf("%d", 322323+8382892392),
		},
		{
			name: "093290592+1=093290593",
			args: args{
				a: "093290592",
				b: "1",
			},
			want: "093290593",
		},
		{
			name: "328023880202+3832827",
			args: args{
				a: "328023880202",
				b: "3832827",
			},
			want: fmt.Sprintf("%d", 328023880202+3832827),
		},
		{
			name: "9938282+1383849",
			args: args{
				a: "9938282",
				b: "1383849",
			},
			want: fmt.Sprintf("%d", 9938282+1383849),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("stringSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightPos(t *testing.T) {
	type args struct {
		a string
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "328904,3->8",
			args: args{
				a: "328904",
				i: 3,
			},
			want: 8,
		},
		{
			name: "8348902,0->2",
			args: args{
				a: "8348902",
				i: 0,
			},
			want: 2,
		},
		{
			name: "8349280,17->0",
			args: args{
				a: "8349280",
				i: 17,
			},
			want: 0,
		},
		{
			name: "4092942848,7->9",
			args: args{
				a: "4092942848",
				i: 7,
			},
			want: 9,
		},
		{
			name: "0,0->0",
			args: args{
				a: "0",
				i: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightPos(tt.args.a, tt.args.i); got != tt.want {
				t.Errorf("rightPos() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestStringMinus(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0-0=0",
			args: args{
				a: "0",
				b: "0",
			},
			want: "0",
		},
		{
			name: "2332-0=2332",
			args: args{
				a: "2332",
				b: "0",
			},
			want: "2332",
		},
		{
			name: "8382892392-322323",
			args: args{
				a: "8382892392",
				b: "322323",
			},
			want: fmt.Sprintf("%d", 8382892392-322323),
		},
		{
			name: "093290592-1=093290591",
			args: args{
				a: "093290592",
				b: "1",
			},
			want: "093290591",
		},
		{
			name: "328023880202-3832827",
			args: args{
				a: "328023880202",
				b: "3832827",
			},
			want: fmt.Sprintf("%d", 328023880202-3832827),
		},
		{
			name: "9938282-1383849",
			args: args{
				a: "9938282",
				b: "1383849",
			},
			want: fmt.Sprintf("%d", 9938282-1383849),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMinus(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("stringSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringByteMultiLinear(t *testing.T){
	stringByteMultiTest(t, byteStringProd{
		name: "stringByteMultiLinear",
		implement: stringByteMultiLinear,
	})
}

func TestStringByteMultiRecursive(t *testing.T){
	stringByteMultiTest(t, byteStringProd{
		name: "stringByteMultiRecursive",
		implement: stringByteMultiRecursive,
	})
}

type byteStringProd struct {
	name string
	implement func (string,byte) string
}

func stringByteMultiTest(t *testing.T, prod byteStringProd) {
	type args struct {
		s string
		m byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"32423424242423*2",
			args: args{
				s: "32423424242423",
				m: '2',
			},
			want: fmt.Sprintf("%d", 32423424242423*2),
		},
		{
			name:"88708372*0",
			args: args{
				s: "88708372",
				m: '0',
			},
			want:"0",
		},
		{
			name:"73273272372393*1",
			args: args{
				s: "73273272372393",
				m: '1',
			},
			want:"73273272372393",
		},
		{
			name:"329282*9",
			args: args{
				s: "329282",
				m: '9',
			},
			want:fmt.Sprintf("%d", 329282*9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prod.implement(tt.args.s, tt.args.m); got != tt.want {
				t.Errorf("%s() = %v, want %v",prod.name, got, tt.want)
			}
		})
	}
}

func BenchmarkStringByteMultiLinear(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		stringByteMultiLinear(a, '7')
	}
}

func BenchmarkStringByteMultiRecursive(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		stringByteMultiRecursive(a, '7')
	}
}