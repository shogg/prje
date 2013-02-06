package e0053

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
	denom := make([]int, max(k, n - k))
	for i := range numer { numer[i] = 1 }
	for i := range denom { denom[i] = 1 }
	for i := 1; i < min(k, n - k); i++ {
		denom[i] = 2
	}

	pnormalize(numer)
	pnormalize(denom)

	for i := range denom {
		m := min(numer[i], denom[i])
		numer[i] -= m
	}

	return mul(numer)
}

func pnormalize(factors []int) {
	for f, e := range factors {
		for _, p := range primefactors(f + 1) {
			factors[p - 1] += e
		}
		factors[f] -= e
	}
}

func primefactors(n int) []int {

	var result []int
	for f := 2; f <= n; f++ {
		for n % f == 0 {
			result = append(result, f)
			n /= f
		}
	}

	return result
}

func min(a, b int) int {
	if a <= b { return a }
	return b
}

func max(a, b int) int {
	if a >= b { return a }
	return b
}

func mul(factors []int) int64 {
	prod := int64(1)
	for f, e := range factors {
		prod *= pow(f + 1, e)
		if prod > 1e18 { break }
	}

	return prod
}

func pow(n, e int) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
	}

	return p
}
