package e0049

import (
	"math"
	"fmt"
	"sort"
)

func E0049() int64 {

	var result []string

	for n := 4321; n <= 6789; n++ {
		d := digits(n, 10)
		if index(d, 0) >= 0 { continue }

		var P []int
		for i := 0; i < 24; i++ {
			p := permutation(i, d)
			if prime(int(number(p))) {
				P = append(P, int(number(p)))
			}
		}

		if len(P) < 3 { continue }

		t := distTriple(P)
		if t == nil { continue }

		r := fmt.Sprintf("%d%d%d", t[0], t[1], t[2])
		if indexS(result, r) < 0 {
			result = append(result, r)
		}
	}

	for _, r := range result {
		println(r)
	}

	return int64(0)
}

func indexS(A []string, e string) int {
	for i, a := range A {
		if a == e { return i }
	}

	return -1
}


func index(A []int, e int) int {
	for i, a := range A {
		if a == e { return i }
	}

	return -1
}

func distTriple(p []int) []int {

	sort.Ints(p)

	for i := 0; i < len(p) - 2; i++ {
		for j := i + 1; j < len(p) - 1; j++ {
			d := p[j] - p[i]
			if d == 0 { continue }
			for k := j + 1; k < len(p); k++ {
				if p[k] == p[j] + d {
					return []int { p[i], p[j], p[k] }
				}
			}
		}
	}

	return nil
}

func permutation(i int, d []int) []int {

	digits := make([]int, len(d))
	copy(digits, d)
	N := len(digits)
	p := make([]int, len(digits))
	f := factorialDigits(i)
	for i := 0; i < len(f); i++ {
		fi := f[N-i-1]
		p[N-i-1] = digits[fi]
		if fi < len(digits) - i - 1 {
			copy(digits[fi:], digits[fi + 1:])
		}
	}

	return p
}

func factorialDigits(n int) []int {

	f := make([]int, 4)
	for i := 1; n != 0; i++ {
		f[i - 1] = n % i
		n /= i
	}

	return f
}

func factorial(n int) int {

	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}

	return f
}

func digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n % base)
		n /= base
	}

	d = append(d, n)

	return d
}

func number(digits []int) int64 {

	n := int64(0)
	for i, d := range digits {
		n += pow(10, i) * int64(d)
	}

	return n
}

func pow(n, e int) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
	}

	return p
}

func prime(p int) bool {
	if p < 2 { return false }
	if p == 2 { return true }
	if p % 2 == 0 { return false }
	for i := 3; i <= int(math.Sqrt(float64(p))); i += 2 {
		if p % i == 0 { return false }
	}

	return true
}
