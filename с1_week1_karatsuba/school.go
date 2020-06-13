package main

import "fmt"

func schoolProduct(a, b string) string {
	if a == "0" || b == "0" || a == "" || b == "" {
		return "0"
	}

	sum := ""
	la := len(a)
	lb := len(b)
	for i := 0; i < la; i++ {
		ai := rightPos(a,i)
		rowP := ""
		rowMem := 0
		for j := 0; j < lb; j++ {
			bj := rightPos(b, j)
			p := ai*bj + rowMem
			ps := fmt.Sprintf("%d", p)
			rowP = ps[len(ps)-1:] + rowP
			rowMem = 0
			if p > 9 {
				rowMem = p/10
			}
		}
		if rowMem > 0 {
			rowP = fmt.Sprintf("%d%s", rowMem, rowP)
		}
		rowP = rowP + nZero(i)
		sum = stringSum(sum, rowP)
	}

	return sum
}
