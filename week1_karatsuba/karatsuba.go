package main

import (
	"strings"
)
func karatsubaProduct(x, y string) string {
	x = strings.TrimLeft(x, "0")
	y = strings.TrimLeft(y, "0")
	lx := len(x)
	ly := len(y)
	if lx == 0 || ly == 0  {
		return "0"
	}
	if lx == 1 {
		if ly == 1 {
			return byteProduct(x[0], y[0])
		}
		return stringByteMultiLinear(y, x[0])
	} else if ly==1 {
		return stringByteMultiLinear(x, y[0])
	}

	if lx < ly {
		x, y, lx, ly = y, x, ly, lx
	}
	diff := lx - ly
	if diff > 0 {
		a := x[:lx-ly]
		b := x[lx-ly:]
		ay := karatsubaProduct(a, y)
		by := karatsubaProduct(b, y)
		ayZero := ay + nZero(ly)
		return stringSum(ayZero, by)
	}

	xMid := lx / 2
	yMid := ly / 2

	a := x[:xMid]
	b := x[xMid:]
	c := y[:yMid]
	d := y[yMid:]

	ac := karatsubaProduct(a, c)
	abSum := stringSum(a, b)
	cdSum := stringSum(c, d)
	abSumCDSum := karatsubaProduct(abSum, cdSum)
	bd := karatsubaProduct(b, d)
	adPlusBc := stringMinus(stringMinus(abSumCDSum, ac), bd)

	acZero := ac + nZero(len(b)+len(d))
	adPlusBcZero := adPlusBc + nZero(len(b))
	prod := stringSum(stringSum(acZero, adPlusBcZero), bd)

	return prod
}
