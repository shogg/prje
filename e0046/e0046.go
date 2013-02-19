package e0046

import (
	"github.com/shogg/prje"
)

func E0046() int64 {

	gold := 0
	for o := 9; ; o++ {

		ok := false
		primes, _ := prje.Psieve(o)
		for _, p := range primes {
			if ok = p == o; ok {
				break
			}

			if ok = prje.Square((o - p) / 2); ok {
				break
			}
		}

		if !ok {
			gold = o;
			break
		}
	}

	return int64(gold)
}
