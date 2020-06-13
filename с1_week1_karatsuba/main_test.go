package main

import "testing"

func BenchmarkKaratsubaProduct(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		karatsubaProduct(a, b)
	}
}

func BenchmarkRecursiveProduct(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		recursiveProduct(a, b)
	}
}

func BenchmarkSchoolProduct(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		schoolProduct(a, b)
	}
}