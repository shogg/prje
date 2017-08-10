package e0058

import (
	"github.com/shogg/prje"
)

func Problem() int64 {

	_, sieve := prje.Psieve(912000000)
	gen := spiralIndexes()

	var c0 int
	var a, b, c, d int
	var s = 3
	for {
		a, b, c, d = gen()

		if !sieve[a] {
			c0++
		}
		if !sieve[b] {
			c0++
		}
		if !sieve[c] {
			c0++
		}
		if !sieve[d] {
			c0++
		}

		if float64(c0)/float64(2*s-1) < .1 {
			break
		}

		s += 2
	}

	return int64(s)
}

func spiralIndexes() func() (int, int, int, int) {

	var i, s = 1, 2
	return func() (int, int, int, int) {
		a, b, c, d := i+1*s, i+2*s, i+3*s, i+4*s
		i += 4 * s
		s += 2
		return a, b, c, d
	}
}
