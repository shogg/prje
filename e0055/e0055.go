package e0055

import (
	"github.com/shogg/prje"
)

func E0055() int64 {

	count := int64(0)
	for n := 10; n < 1e5; n++ {
		m := int64(n)
		lychrel := true
		for i := 0; i < 50; i++ {
			if n == 196 { println(m) }
			m += reverse(m)
			if prje.Pal(prje.Digits(int(m), 10)) {
				lychrel = false
				break
			}
		}

		if lychrel { count++ }
	}

	return count
}

func reverse(n int64) int64 {
	r := prje.Number(prje.Reverse(prje.Digits(int(n), 10)), 10)
	return r
}
