package e0041

import (
	"github.com/shogg/prje"
)

func E0041() int64 {

	digits := []int { 7, 6, 5, 4, 3, 2, 1 }
	n := len(digits)

	p := int64(0)
	for i := 0; int64(i) < prje.Factorial(n); i++ {
		p = prje.Number(prje.Permutation(i, digits), 10)
		if prje.Prime(int(p)) { break }
	}

	return p
}
