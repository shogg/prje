package e0032

import (
	"sort"
	"github.com/shogg/prje"
)

func E0032() int64 {

	var numbers []int
	for a := 1; a < 9998; a++ {
		for b := a + 1; b < 9999; b++ {
			d := append(prje.Digits(a, 10), prje.Digits(b, 10)...)
			d = append(d, prje.Digits(a*b, 10)...)
			if len(d) > 9 { break }
			if prje.Pandigital(9, d) {
				numbers = append(numbers, a*b)
			}
		}
	}

	sort.Ints(numbers)
	uniq := prje.Uniq(numbers)

	sum := 0
	for _, n := range uniq {
		sum += n
	}

	return int64(sum)
}
