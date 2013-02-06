package e0036

import (
	. "github.com/shogg/prje/lib"
)

func E0036() int64 {

	sum := 0
	for n := 1; n < 1e6; n += 2 {
		if Pal(Digits(n, 2)) &&
		   Pal(Digits(n, 10)) {
			sum += n
		}
	}

	return int64(sum)
}
