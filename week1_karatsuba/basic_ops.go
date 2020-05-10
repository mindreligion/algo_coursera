package main

import (
	"fmt"
	"log"
)

func byteProduct(a byte, b byte) string {
	x := byteToInt(a)
	y := byteToInt(b)
	return fmt.Sprintf("%d", x*y)
}

func byteToInt(a byte) int {
	if a < 48 || 57 < a {
		log.Fatalf("non numeric byte : %v", a)
	}
	return int(a) - 48
}

func stringSum(a, b string) string {
	la, lb := len(a), len(b)
	l := la
	if lb > la {
		l = lb
	}

	sum := ""
	mem := 0
	for i := 0; i < l; i++ {
		s := rightPos(a, i) + rightPos(b, i) + mem
		mem = 0
		ss := fmt.Sprintf("%d", s)
		ls := len(ss)
		sum = ss[ls-1:] + sum
		if s > 9 {
			mem = s / 10
		}
	}

	if mem > 0 {
		sum = fmt.Sprintf("%d", mem) + sum
	}

	return sum
}

func rightPos(a string, i int) int {
	l := len(a)
	if i >= l {
		return 0
	}

	return byteToInt(a[l-1-i])
}

func nZero(n int) string {
	zeros := ""
	for i := 0; i < n; i++ {
		zeros += "0"
	}
	return zeros
}

func stringMinus(a, b string) string {
	l := len(a)

	sum := ""
	mem := 0
	for i := 0; i < l; i++ {
		s := rightPos(a, i) - rightPos(b, i) + mem
		mem = 0
		if s < 0 {
			s += 10
			mem = -1
		}
		ss := fmt.Sprintf("%d", s)
		sum = ss + sum
	}

	if mem > 0 {
		sum = fmt.Sprintf("%d", mem) + sum
	}

	return sum
}

func stringByteMultiLinear(s string, mByte byte) string {
	m := byteToInt(mByte)
	if m == 0 {
		return "0"
	}
	l := len(s)
	mem := 0
	prod := ""
	for j := 0; j < l; j++ {
		sj := rightPos(s, j)
		p := m*sj + mem
		ps := fmt.Sprintf("%d", p)
		prod = ps[len(ps)-1:] + prod
		mem = 0
		if p > 9 {
			mem = p/10
		}
	}
	if mem >0 {
		prod = fmt.Sprintf("%d%s", mem, prod)
	}

	return prod
}

func stringByteMultiRecursive(s string, mByte byte) string {
	m := byteToInt(mByte)
	if m == 0 {
		return "0"
	}
	l := len(s)
	if len(s) == 1 {
		return fmt.Sprintf("%d", byteToInt(s[0]) * m)
	}
	mid := l/2
	return stringSum(stringByteMultiRecursive(s[:mid], mByte) + nZero(l-mid), stringByteMultiRecursive(s[mid:], mByte))
}
