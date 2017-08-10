package prje

import (
	"math"
	"sort"
)

func Digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n%base)
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

func PowMod(n, e int, mod int64) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
		p %= mod
	}

	return p
}

func Reverse(digits []int) []int {
	r := make([]int, len(digits))
	copy(r, digits)

	size := len(r)
	for i := range r {
		r[i], r[size-i-1] = r[size-i-1], r[i]
		if i >= size-i-2 {
			break
		}
	}

	return r
}

func Pal(digits []int) bool {

	dlen := len(digits)
	for i := 0; i < dlen/2; i++ {
		if digits[i] != digits[dlen-i-1] {
			return false
		}
	}

	return true
}

func Psieve(n int) ([]int, []bool) {

	var primes = make([]int, 0, n+1)
	var sieve = make([]bool, n+1)

	for p := 2; p <= n; p++ {
		if sieve[p] {
			continue
		}

		for i := 2 * p; i <= n; i += p {
			sieve[i] = true
		}

		primes = append(primes, p)
	}

	return primes, sieve
}

func Prime(p int) bool {
	if p < 2 {
		return false
	}
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(p))); i += 2 {
		if p%i == 0 {
			return false
		}
	}

	return true
}

func Primefactors(n int) []int {

	var result []int
	for f := 2; f <= n; f++ {
		for n%f == 0 {
			result = append(result, f)
			n /= f
		}
	}

	return result
}

func Factors(n int) []int {

	result := []int{1}
	sqrt := int(math.Sqrt((float64(n))))
	for f := 2; f <= sqrt; f++ {
		if n%f == 0 {
			result = append(result, f)
			j := n / f
			if j != f {
				result = append(result, j)
			}
		}
	}

	return result
}

func Gcd(a, b int) int {

	min := a
	if b < a {
		min = b
	}

	for n := min; n >= 1; n++ {
		if a%n == 0 && b%n == 0 {
			return n
		}
	}

	return 1
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Factorial(n int) int64 {

	f := int64(1)
	for i := 1; i <= n; i++ {
		f *= int64(i)
	}

	return f
}

func Pandigital(n int, digits []int) bool {
	if len(digits) != n {
		return false
	}

	tmp := make([]int, len(digits))
	copy(tmp, digits)

	sort.Ints(tmp)

	for i, d := range tmp {
		if d != i+1 {
			return false
		}
	}

	return true
}

func Permutation(i int, d []int) []int {

	N := len(d)
	digits := make([]int, N)
	copy(digits, d)
	p := make([]int, N)
	f := FactorialDigits(i, N)
	for i := 0; i < len(f); i++ {
		fi := f[N-i-1]
		p[N-i-1] = digits[fi]
		if fi < len(digits)-i-1 {
			copy(digits[fi:], digits[fi+1:])
		}
	}

	return p
}

func FactorialDigits(n, count int) []int {

	f := make([]int, count)
	for i := 1; n != 0; i++ {
		f[i-1] = n % i
		n /= i
	}

	return f
}

func Uniq(numbers []int) []int {
	var uniq []int
	last := math.MinInt32
	for _, n := range numbers {
		if n == last {
			continue
		}
		uniq = append(uniq, n)
		last = n
	}

	return uniq
}

func Square(n int) bool {
	r := math.Sqrt(float64(n))
	return r == math.Floor(r)
}

func Abundant(n int) bool {
	sum := 0
	for _, d := range Factors(n) {
		sum += d
	}

	return sum > n
}

func Triangle(n int64) int64 {
	t := n * (n + 1) / 2
	return t
}

func Triangular(t int) bool {
	n := -.5 + math.Sqrt(.25+2.0*float64(t))
	return n == math.Trunc(n)
}

func Pentagon(n int64) int64 {
	return n * (3*n - 1) / 2
}

func Pentagonal(p int64) bool {
	n := 1.0/6.0 + math.Sqrt(1.0/36.0+2.0*float64(p)/3.0)
	return Pentagon(int64(n)) == p
}

func Hexagonal(h int64) bool {
	n := .25 + math.Sqrt(1.0/16.0+float64(h)/2.0)
	return n == math.Trunc(n)
}
