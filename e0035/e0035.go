package e0035

import (
	"math"
)

func E0035() int64 {

	count := 0
	for n := 2; n < 1000000; n++ {
		d := digits(n)
		if !prime(n) { continue }

		cprime := true
		for i := 1; i < len(d); i++ {
			circle(d)
			if !prime(number(d)) {
				cprime = false; break
			}
		}

		if cprime {
			count++
		}
	}

	return int64(count)
}

func circle(digits []int) {

	tmp := digits[0]
	for i := 0; i < len(digits) - 1; i++ {
		digits[i] = digits[i + 1]
	}
	digits[len(digits) - 1] = tmp
}

func digits(n int) []int {

	var d []int
	for n >= 10 {
		d = append(d, n % 10)
		n /= 10
	}

	d = append(d, n)

	return d
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
