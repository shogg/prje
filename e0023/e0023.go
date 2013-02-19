package e0023

import (
	"github.com/shogg/prje"
)

func E0023() int64 {

	sum := 0
	for n := 1; n < 28123; n++ {
		found := true
		for a := 1; a < n; a++ {
			if prje.Abundant(a) && prje.Abundant(n - a) {
				found = false; break
			}
		}
		if found { sum += n }
	}

	return int64(sum)
}
