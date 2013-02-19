package e0045

import (
	"github.com/shogg/prje"
)

func E0045() int64 {

	result := int64(0)
	for n := int64(285+1); ; n++ {
		t := prje.Triangle(n)
		if prje.Pentagonal(t) && prje.Hexagonal(t) {
			result = t; break
		}
	}

	return result
}
