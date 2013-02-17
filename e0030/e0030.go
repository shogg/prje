package e0030

import (
	"github.com/shogg/prje"
)

func E0030() int64 {

	var numbers []int

	for n := 10; n <= 5 * int(prje.Pow(9, 5)); n++ {
		psum := 0
		for _, d := range prje.Digits(n, 10) {
			psum += int(prje.Pow(d, 5))
		}

		if psum == n {
			numbers = append(numbers, n)
		}
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return int64(sum)
}
