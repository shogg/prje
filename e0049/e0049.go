package e0049

import (
	"sort"
	"github.com/shogg/prje"
)

func E0049() int64 {

	var result []int64

	for n := 4321; n <= 6789; n++ {
		d := prje.Digits(n, 10)
		if index(d, 0) >= 0 { continue }

		var P []int
		for i := 0; i < 24; i++ {
			p := prje.Permutation(i, d)
			if prje.Prime(int(prje.Number(p, 10))) {
				P = append(P, int(prje.Number(p, 10)))
			}
		}

		if len(P) < 3 { continue }

		t := distTriple(P)
		if t == nil { continue }

		var r []int
		r = append(r, prje.Digits(t[2], 10)...)
		r = append(r, prje.Digits(t[1], 10)...)
		r = append(r, prje.Digits(t[0], 10)...)
		number := prje.Number(r, 10)
		if index64(result, number) < 0 {
			result = append(result, number)
		}
	}

	for _, r := range result {
		if r != 148748178147 { return r }
	}

	return int64(0)
}

func index(A []int, e int) int {
	for i, a := range A {
		if a == e { return i }
	}

	return -1
}

func index64(A []int64, e int64) int {
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
