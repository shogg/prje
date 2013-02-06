package e0023

import (
	"math"
)

func E0023() int64 {

	sum := 0
	for n := 1; n < 28123; n++ {
		found := true
		for a := 1; a < n; a++ {
			if abundant(a) && abundant(n - a) { found = false; break }
		}
		if found { sum += n }
	}

	return int64(sum)
}

func abundant(n int) bool {

	sum := 0
	for _, d := range divisors(n) {
		sum += d
	}

	return sum > n
}

func divisors(n int) []int {
	d := []int { 1 }

	sqrt := int(math.Sqrt((float64(n))))
	for i := 2; i <= sqrt; i++ {
		if n % i == 0 {
			d = append(d, i)
			j := n / i
			if j != i { d = append(d, j) }
		}
	}

	return d
}
