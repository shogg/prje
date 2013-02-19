package e0044

import (
	"github.com/shogg/prje"
)

func E0044() int64 {

	d := int64(0)
	for i := int64(2); d == 0 ; i++ {
		for j := int64(i - 1); j > 0; j-- {
			p1 := prje.Pentagon(i) + prje.Pentagon(j)
			p2 := prje.Pentagon(i) - prje.Pentagon(j)
			if prje.Pentagonal(p1) && prje.Pentagonal(p2) {
				d = p2; break
			}
		}
	}

	return d
}
