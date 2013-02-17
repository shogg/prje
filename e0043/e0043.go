package e0043

import (
	"github.com/shogg/prje"
)

func E0043() int64 {

	N := 10
	digits := []int { 9, 8, 7, 6, 5, 4, 3, 2, 1, 0 }
	primes := []int { 2, 3, 5, 7, 11, 13, 17 }

	sum := int64(0)
	for i := 0; i < int(prje.Factorial(N)); i++ {
		p := prje.Permutation(i, digits)

		ok := true
		for j := 0; j < len(primes); j++ {
			n3 := prje.Number(p[N-j-4:N-j-1], 10)
			prime := int64(primes[j])
			ok = n3 % prime == 0

			if !ok { break }
		}

		if ok {
			sum += prje.Number(p, 10)
		}
	}

	return sum
}
