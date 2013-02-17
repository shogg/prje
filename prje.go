package prje

import (
	"math"
	"sort"
)

func Digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n % base)
		n /= base
	}

	d = append(d, n)

	return d
}

func Number(digits []int, base int) int64 {

	n := int64(0)
	for i, d := range digits {
		n += Pow(base, i) * int64(d)
	}

	return n
}

func Pow(n, e int) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
	}

	return p
}

func Reverse(digits []int) []int {
	r := make([]int, len(digits))
	copy(r, digits)

	size := len(r)
	for i := range r {
		r[i], r[size - i - 1] = r[size - i - 1], r[i]
		if i >= size - i - 2 { break }
	}

	return r
}

func Pal(digits []int) bool {

	dlen := len(digits)
	for i := 0; i < dlen / 2; i++ {
		if digits[i] != digits[dlen - i - 1] {
			return false
		}
	}

	return true
}

func Psieve(n int) ([]int, []bool) {

	var primes []int
	var sieve = make([]bool, n + 1)

	for p := 2; p <= n; p++ {
		if sieve[p] { continue }

		for i := 2*p; i <= n; i += p {
			sieve[i] = true
		}

		primes = append(primes, p)
	}

	return primes, sieve
}

func Prime(p int) bool {
	if p < 2 { return false }
	if p == 2 { return true }
	if p % 2 == 0 { return false }
	for i := 3; i <= int(math.Sqrt(float64(p))); i += 2 {
		if p % i == 0 { return false }
	}

	return true
}

func Gcd(a, b int) int {

	min := a
	if b < a { min = b }

	for n := min; n >= 1; n++ {
		if a % n == 0 && b % n == 0 {
			return n
		}
	}

	return 1
}

func Factorial(n int) int64 {

	f := int64(1)
	for i := 1; i <= n; i++ {
		f *= int64(i)
	}

	return f
}

func Pandigital(n int, digits []int) bool {
	if len(digits) != n { return false }

	tmp := make([]int, len(digits))
	copy(tmp, digits)

	sort.Ints(tmp)

	for i, d := range tmp {
		if d != i + 1 { return false }
	}

	return true
}

func Permutation(i int, d []int) []int {

	digits := make([]int, len(d))
	copy(digits, d)
	N := len(digits)
	p := make([]int, len(digits))
	f := FactorialDigits(i)
	for i := 0; i < len(f); i++ {
		fi := f[N-i-1]
		p[N-i-1] = digits[fi]
		if fi < len(digits) - i - 1 {
			copy(digits[fi:], digits[fi + 1:])
		}
	}

	return p
}

func FactorialDigits(n int) []int {

	f := make([]int, 4)
	for i := 1; n != 0; i++ {
		f[i - 1] = n % i
		n /= i
	}

	return f
}
