package e0038

import (
	"sort"
)

func E0038() int64 {
	max := 0
	for n := 9876; n > 1; n-- {
		var d []int
		for i := 1; i <= 4; i++ {
			d = append(digits(i * n, 10), d...)
			if len(d) > 9 { break }
			if pandigital(9, d) {
				p := number(d)
				if p > max {
					max = p
				}
			}
		}
	}
	return int64(max)
}

func pandigital(n int, digits []int) bool {
	if len(digits) != n { return false }

	tmp := make([]int, len(digits))
	copy(tmp, digits)

	sort.Ints(tmp)

	for i, d := range tmp {
		if d != i + 1 { return false }
	}

	return true
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
