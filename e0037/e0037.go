package e0037

import (
	"math"
)

func E0037() int64 {

	sum := 0
	count := 0
	for n := 11; count < 11; n++ {
		d := digits(n)
		if !prime(n) { continue }

		tprime := true
		for i := 1; i < len(d); i++ {
			if !prime(number(d[i:])) ||
			   !prime(number(d[:i])) {
				tprime = false; break
			}
		}

		if tprime {
			sum += n
			count++
		}
	}

	return int64(sum)
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
