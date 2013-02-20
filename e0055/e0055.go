package e0055

import (
	"github.com/shogg/prje"
)

func E0055() int64 {

	count := int64(0)
	for n := 10; n < 1e4; n++ {
		m := prje.NewNat(prje.Digits(n, 10), 10)
		lychrel := true
		for i := 0; i < 50; i++ {
			m = m.Plus(reverse(m))
			if prje.Pal(m.Digits) {
				lychrel = false
				break
			}
		}

		if lychrel { count++ }
	}

	return count
}

func reverse(n prje.Nat) prje.Nat {
	r := prje.NewNat(prje.Reverse(n.Digits), 10)
	return r
}
