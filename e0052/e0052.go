package e0052

import (
	"sort"
	"github.com/shogg/prje"
)

func E0052() int64 {

	n := 1
	for ; ; n++ {
		d1 := prje.Digits(n, 10)
		sort.Ints(d1)
		found := true
		for x := 2; x <= 6; x++ {
			dx := prje.Digits(n*x, 10)
			sort.Ints(dx)
			if !sameset(d1, dx) {
				found = false; break
			}
		}
		if found { break }
	}

	return int64(n)
}

func sameset(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] { return false }
	}
	return true
}
