package main

import (
	"fmt"
)

func main() {
	count := 0
	for n := 1; n <= 100; n++ {
		for k := 1; k <= n / 2; k++ {
			if C(n, k) > 1e6 {
				count += 2*(n / 2 - k)
				break
			}
		}
	}

	println(count)
}

func C(n, k int) int64 {

	// n!/(k!(n-k)!)

	numer := make([]int, n)
	m := max(k, n - k)
	denom := make([]int, m)
	for i := range numer { numer[i] = 1 }
	for i := range denom { denom[i] = 1 }
	for i := 1; i < min(k, n - k); i++ {
		denom[i] = 2
	}

	pnormalize(numer)
	if n == 5 && k == 2 {
		fmt.Println(numer)
	}
	pnormalize(denom)

	for i := range denom {
		m := min(numer[i], denom[i])
		numer[i] -= m
		denom[i] -= m
	}

	return mul(numer) / mul(denom)
}

func pnormalize(factors []int) {
	for f := range factors {
		for _, p := range primefactors(f) {
			factors[p]++
		}
		factors[f]--
	}
}

func primefactors(n int) []int {

	var result []int
	for f := 2; f <= n; f++ {
		if n % f == 0 {
			result = append(result, f)
		}
		for n % f == 0 {
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
		prod *= pow(f, e)
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
