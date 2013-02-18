package e0021

import (
	"github.com/shogg/prje"
)

func E0021() int64 {

	sum := 0
	for a := 1; a < 10000; a++ {
		if b := d(a); d(b) == a && a != b {
			sum += a
		}
	}

	return int64(sum)
}

func d(n int) int {

	sum := 0
	for _, f := range prje.Factors(n) {
		sum += f
	}

	return sum
}
