package e0060

import (
	"strconv"

	"github.com/shogg/prje"
)

func Problem() int64 {

	p4 := []int{3, 7, 109, 673}

	primes, sieve := prje.Psieve(1e6)

	for _, p := range primes {

		if contains(p4, p) {
			continue
		}

		if !primeCat(sieve, p, p4) {
			continue
		}

		return int64(sum(p4...) + p)
	}

	return -1
}

func primeCat(sieve []bool, p0 int, primes []int) bool {

	cat := func(p0, p1 int) int {
		s0 := strconv.Itoa(p0)
		s1 := strconv.Itoa(p1)
		p2, _ := strconv.Atoi(s0 + s1)
		return p2
	}

	for _, p := range primes {

		cat0 := cat(p0, p)
		cat1 := cat(p, p0)

		if cat0 >= len(sieve) || cat1 >= len(sieve) {
			panic("sieve too small")
		}

		if sieve[cat0] || sieve[cat1] {
			return false
		}
	}

	return true
}

func contains(xs []int, n int) bool {
	for _, x := range xs {
		if x == n {
			return true
		}
	}
	return false
}

func sum(xs ...int) int {

	var s int
	for _, x := range xs {
		s += x
	}

	return s
}
