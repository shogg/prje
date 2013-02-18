package e0053

import (
	"github.com/shogg/prje"
)

func E0053() int64 {
	count := 0
	for n := 1; n <= 100; n++ {
		for k := 1; k <= n / 2; k++ {
			if C(n, k) > 1e6 {
				count += n - 2*(k - n % 2)
				break
			}
		}
	}

	return int64(count)
}

// n!/(k!(n - k)!)
func C(n, k int) int64 {

	numer := make([]int, n)
	denom := make([]int, prje.Max(k, n - k))
	for i := range numer { numer[i] = 1 }
	for i := range denom { denom[i] = 1 }
	for i := 1; i < prje.Min(k, n - k); i++ {
		denom[i] = 2
	}

	pnormalize(numer)
	pnormalize(denom)

	for i := range denom {
		m := prje.Min(numer[i], denom[i])
		numer[i] -= m
	}

	return mul(numer)
}

func pnormalize(factors []int) {
	for f, e := range factors {
		for _, p := range prje.Primefactors(f + 1) {
			factors[p - 1] += e
		}
		factors[f] -= e
	}
}

func mul(factors []int) int64 {
	prod := int64(1)
	for f, e := range factors {
		prod *= prje.Pow(f + 1, e)
		if prod > 1e18 { break }
	}

	return prod
}
