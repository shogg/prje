package e0037

import (
	"github.com/shogg/prje"
)

func E0037() int64 {

	sum := 0
	count := 0
	for n := 11; count < 11; n++ {
		d := prje.Digits(n, 10)
		if !prje.Prime(n) { continue }

		tprime := true
		for i := 1; i < len(d); i++ {
			if !prje.Prime(int(prje.Number(d[i:], 10))) ||
			   !prje.Prime(int(prje.Number(d[:i], 10))) {
				tprime = false; break
			}
		}

		if tprime {
			sum += n
			count++
		}
	}

	return int64(sum)
}
