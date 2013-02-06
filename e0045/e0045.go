package e0045

import (
	"math"
)

func E0045() int64 {

	result := int64(0)
	for n := int64(285+1); ; n++ {
		t := triangle(n)
		if pentagonal(t) && hexagonal(t) {
			result = t; break
		}
	}

	return result
}

func triangle(n int64) int64 {
	t := n*(n + 1) / 2
	return t
}

func pentagonal(p int64) bool {
	n := 1.0/6.0 + math.Sqrt(1.0/36.0 + 2.0/3.0*float64(p))
	return n == math.Trunc(n)
}

func hexagonal(h int64) bool {
	n := .25 + math.Sqrt(1.0/16.0 + float64(h)/2.0)
	return n == math.Trunc(n)
}
