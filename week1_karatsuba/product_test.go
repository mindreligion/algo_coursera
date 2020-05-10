package main

import (
	"fmt"
	"testing"
)

type product struct {
	algo string
	implementation func(a,b string) string
}

func TestKaratsubaProduct(t *testing.T) {
	testProduct(t, product {
		algo: "karatsubaProduct",
		implementation: karatsubaProduct,
	})
}

func TestRecursiveProduct(t *testing.T) {
	testProduct(t, product {
		algo: "recursiveProduct",
		implementation: recursiveProduct,
	})
}

func TestSchoolProduct(t *testing.T) {
	testProduct(t, product {
		algo: "schoolProduct",
		implementation: schoolProduct,
	})
}

func testProduct(t *testing.T, prod product) {
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
			name: "24*32",
			args: args{
				a: "24",
				b: "32",
			},
			want: fmt.Sprintf("%d", 24*32),
		},
		{
			name: "39*4234252532",
			args: args{
				a: "39",
				b: "4234252532",
			},
			want: fmt.Sprintf("%d", 39*4234252532),
		},
		{
			name: "9032457*23458",
			args: args{
				a: "9032457",
				b: "23458",
			},
			want: fmt.Sprintf("%d", 9032457*23458),
		},
		{
			name: "893923827*893923827",
			args: args{
				a: "893923827",
				b: "893923827",
			},
			want: fmt.Sprintf("%d", int64(893923827*893923827)),
		},
		{
			name: "0*32234585",
			args: args{
				a: "0",
				b: "32234585",
			},
			want: fmt.Sprintf("%d", 0),
		},
		{
			name: "0*0",
			args: args{
				a: "0",
				b: "0",
			},
			want: fmt.Sprintf("%d", 0),
		},
		{
			name: "23758852*32835829",
			args: args{
				a: "23758852",
				b: "32835829",
			},
			want: fmt.Sprintf("%d", 23758852*32835829),
		},
		{
			name: "93280238*8393283285",
			args: args{
				a: "93280238",
				b: "8393283285",
			},
			want: fmt.Sprintf("%d", int64(93280238)*int64(8393283285)),
		},
		{
			name: "323223*53432",
			args: args{
				a: "323223",
				b: "53432",
			},
			want: fmt.Sprintf("%d", 323223*53432),
		},
		{
			name: "12345678998765432112345678987654321*1212121212121234343434343434343434444444444",
			args: args{
				a: "12345678998765432112345678987654321",
				b: "1212121212121234343434343434343434444444444",
			},
			want: "14964459392443222363386654166355210140914770544956011896745221548821561042524",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prod.implementation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("%s() = %v, want %v", prod.algo, got, tt.want)
			}
		})
	}
}
