package e0036

import (
	"github.com/shogg/prje"
)

func E0036() int64 {

	sum := 0
	for n := 1; n < 1e6; n += 2 {
		if prje.Pal(prje.Digits(n, 2)) &&
		   prje.Pal(prje.Digits(n, 10)) {
			sum += n
		}
	}

	return int64(sum)
}
