package prje

// n!/(k!(n - k)!)
func C(n, k int) int64 {

	numer := make([]int, n)
	denom := make([]int, Max(k, n - k))
	for i := range numer { numer[i] = 1 }
	for i := range denom { denom[i] = 1 }
	for i := 1; i < Min(k, n - k); i++ {
		denom[i] = 2
	}

	pnormalize(numer)
	pnormalize(denom)

	for i := range denom {
		m := Min(numer[i], denom[i])
		numer[i] -= m
	}

	return mul(numer)
}

func pnormalize(factors []int) {
	for f, e := range factors {
		for _, p := range Primefactors(f + 1) {
			factors[p - 1] += e
		}
		factors[f] -= e
	}
}

func mul(factors []int) int64 {
	prod := int64(1)
	for f, e := range factors {
		prod *= Pow(f + 1, e)
		if prod > 1e18 { break }
	}

	return prod
}
