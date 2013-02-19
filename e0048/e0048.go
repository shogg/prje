package e0048

import (
	"github.com/shogg/prje"
)

func E0048() int64 {

	sum := int64(0)
	for a := 1; a <= 1e3; a++ {
		sum += prje.PowMod(a, a, 1e10)
	}

	sum %= 1e10
	return sum
}
