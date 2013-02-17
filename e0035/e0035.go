package e0035

import (
	"github.com/shogg/prje"
)

func E0035() int64 {

	count := 0
	for n := 2; n < 1000000; n++ {
		d := prje.Digits(n, 10)
		if !prje.Prime(n) { continue }

		cprime := true
		for i := 1; i < len(d); i++ {
			circle(d)
			if !prje.Prime(int(prje.Number(d, 10))) {
				cprime = false; break
			}
		}

		if cprime {
			count++
		}
	}

	return int64(count)
}

func circle(digits []int) {

	tmp := digits[0]
	for i := 0; i < len(digits) - 1; i++ {
		digits[i] = digits[i + 1]
	}
	digits[len(digits) - 1] = tmp
}
