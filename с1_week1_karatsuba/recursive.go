package main

import (
	"strings"
)

func recursiveProduct (x, y string) string {
	x = strings.TrimLeft(x,"0")
	y = strings.TrimLeft(y,"0")
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

	xMid := lx /2
	yMid := ly /2

	a := x[:xMid]
	b := x[xMid:]
	c := y[:yMid]
	d := y[yMid:]


	ac := recursiveProduct(a,c)
	ad := recursiveProduct(a,d)
	bc := recursiveProduct(b, c)
	bd := recursiveProduct(b,d)

	padB := nZero(len(b))
	padD := nZero(len(d))

	acZero := ac + nZero(len(b)+len(d))
	adZero := ad + padB
	bcZero := bc + padD
	prod := stringSum(stringSum(stringSum(acZero,adZero),bcZero),bd)
	return prod
}

