package e0010

import (
	"github.com/shogg/prje"
)

func E0010() int64 {

	primes, _ := prje.Psieve(2000000)

	sum := int64(0)
	for _, p := range primes {
		sum += int64(p)
	}

	return sum
}
