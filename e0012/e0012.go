package e0012

import (
	"math"
)

func E0012() int64 {

	triangle := int64(0)
	for i := int64(1); divisors(triangle) <= 500; i++ {
		triangle += i
	}

	return triangle
}

func divisors(n int64) int {

	d := 0
	sqrt := int64(math.Sqrt((float64(n))))
	for i := int64(1); i <= sqrt; i++ {
		if n % i == 0 { d += 2 }
	}

	if sqrt*sqrt == n { d-- }

	return d
}
