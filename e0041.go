package main

import (
	"math"
)

const (
	N = 7
)

func main() {

	digits := [...]int { 7, 6, 5, 4, 3, 2, 1 }

	p := 0
	for i := 0; i < factorial(N); i++ {
		p = number(permutation(i, digits))
		if prime(p) { break }
	}

	println(p)
}

func permutation(i int, digits [N]int) []int {

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

	f := make([]int, N)
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

func number(digits []int) int {

	n := 0
	for i, d := range digits {
		n += pow(10, i) * d
	}

	return n
}

func pow(n, e int) int {

	p := 1
	for i := 1; i <= e; i++ {
		p *= n
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
