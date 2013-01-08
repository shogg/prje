package main

import (
	"math"
)

func main() {

	gold := 0
	for o := 9; ; o++ {

		ok := false
		for _, p := range primes(o) {
			if ok = p == o; ok {
				break
			}

			if ok = square((o - p) / 2); ok {
				break
			}
		}

		if !ok {
			gold = o;
			break
		}
	}

	println(gold)
}

func primes(n int) []int {

	var primes []int

	sieve := make([]bool, n + 1)
	for p := 2; p <= n; p++ {
		if sieve[p] { continue }

		for i := 2*p; i <= n; i += p {
			sieve[i] = true
		}

		primes = append(primes, p)
	}

	return primes
}

func square(n int) bool {
		r := math.Sqrt(float64(n))
		return r == math.Floor(r)
}
