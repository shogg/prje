package e0050

import (
	"github.com/shogg/prje"
)

const N = 1e6

func E0050() int64 {

	primes, sieve := prje.Psieve(N)

	max := 0
	prime := 0

	for i, _ := range primes {
		sum := 0
		count := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			count++

			if sum > N { break }

			if !sieve[sum] && count > max {
				prime = sum
				max = count
			}
		}
	}

	return int64(prime)
}
