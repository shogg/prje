package e0051

import (
	"github.com/shogg/prje"
)

func E0051() int64 {

	result := 0
	primes, sieve := prje.Psieve(999999)

out:
	for _, p := range primes {
		d := prje.Digits(p, 10)
		for min := 0; min <= 2; min++ {

			idx := indexes(d, min)
			for m := 1; log2(m) + 1 <= len(idx); m++ {

				count := 1
				d1 := make([]int, len(d)); copy(d1, d)
				for i := min; i < 9; i++ {
					n := prje.Number(inc(d1, mask(idx, m)), 10)
					if sieve[n] == false {
						count++
					}
				}

				if count >= 8 {
					result = p; break out
				}
			}
		}
	}

	return int64(result)
}

func inc(digits, idx []int) []int {
	for _, i := range idx {
		digits[i]++
	}
	return digits
}

func mask(idx []int, m int) []int {

	var masked []int
	for _, i := range idx {
		if m & 0x1 == 1 {
			masked = append(masked, i)
		}
		m >>= 1
	}

	return masked
}

func indexes(digits []int, d int) []int {

	var idx []int
	for i, n := range digits {
		if n == d {
			idx = append(idx, i)
		}
	}

	return idx
}

func log2(n int) int {
	log := 0
	for ; n >= 2; n >>= 1 {
		log++
	}
	return log
}
