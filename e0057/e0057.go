package e0057

import (
	"math/big"
)

func E0057() int64 {

	zero := big.NewRat(int64(0), int64(1))
	one := big.NewRat(int64(1), int64(1))
	two := big.NewRat(int64(2), int64(1))
	
	v, r := &big.Rat{}, &big.Rat{}
	r.Set(zero)

	count := 0
	for i := 0; i < 1000; i++ {
		
		r.Quo(one, r.Add(two, r))
		v.Add(one, r)

		if len(v.Num().String()) > len(v.Denom().String()) {
			count++
		}
	}

	return int64(count)
}
