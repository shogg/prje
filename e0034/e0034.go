package e0034

import (
	"github.com/shogg/prje"
)

func E0034() int64 {

	sum := int64(0)
	for n := 3; n < 8 * int(prje.Factorial(9)); n++ {

		s := int64(0)
		for _, d := range prje.Digits(n, 10) {
			s += prje.Factorial(d)
		}

		if s == int64(n) {
			sum += s
		}

	}

	return sum
}
