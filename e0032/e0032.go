package e0032

import (
	"sort"
)

func E0032() int64 {

	var numbers []int
	for a := 1; a < 9998; a++ {
		for b := a + 1; b < 9999; b++ {
			d := append(append(digits(a), digits(b)...), digits(a*b)...)
			if len(d) > 9 { break }
			if pandigital(9, d) {
				numbers = append(numbers, a*b)
			}
		}
	}

	sort.Ints(numbers)

	var uniq []int
	last := 0
	for _, n := range numbers {
		if n == last { continue }
		uniq = append(uniq, n)
		last = n
	}

	sum := 0
	for _, n := range uniq {
		sum += n
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

func pandigital(n int, digits []int) bool {
	if len(digits) != n { return false }

	sort.Ints(digits)

	for i, d := range digits {
		if d != i + 1 { return false }
	}

	return true
}
