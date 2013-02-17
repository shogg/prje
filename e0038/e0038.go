package e0038

import (
	"github.com/shogg/prje"
)

func E0038() int64 {
	max := int64(0)
	for n := 9876; n > 1; n-- {
		var d []int
		for i := 1; i <= 4; i++ {
			d = append(prje.Digits(i * n, 10), d...)
			if len(d) > 9 { break }
			if prje.Pandigital(9, d) {
				p := prje.Number(d, 10)
				if p > max {
					max = p
				}
			}
		}
	}
	return max
}
