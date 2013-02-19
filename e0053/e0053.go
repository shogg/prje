package e0053

import (
	"github.com/shogg/prje"
)

func E0053() int64 {
	count := 0
	for n := 1; n <= 100; n++ {
		for k := 1; k <= n / 2; k++ {
			if prje.C(n, k) > 1e6 {
				count += n - 2*(k - n % 2)
				break
			}
		}
	}

	return int64(count)
}
