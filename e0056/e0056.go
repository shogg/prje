package e0056

import (
	"math/big"
	"strconv"
)

func E0056() int64 {

	var max int
	for i := int64(1); i < 100; i++ {
		for j := int64(1); j < 100; j++ {
			a, b := big.NewInt(i), big.NewInt(j)
			c := a.Exp(a, b, nil)

			sum := sumDigits(c)
			if sum > max {
				max = sum
				// fmt.Printf("%d**%d = %d\n", i, j, max)
			}
		}
	}

	return int64(max)
}

func sumDigits(v *big.Int) int {
	var sum int
	for _, s := range v.String() {
		d, _ := strconv.Atoi(string(s))
		sum += d
	}

	return sum
}
