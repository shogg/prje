package e0025

import (
	"math/big"
)

func E0025() int64 {

	e999 := big.NewInt(10)
	e999 = e999.Exp(e999, big.NewInt(999), nil)

	fib := make(chan *big.Int)
	go fibonacci(fib)

	i := 1
	for f := range fib {
		if f.Cmp(e999) >= 0 { break }
		i++
	}

	return int64(i)
}

func fibonacci(fib chan *big.Int) {

	f1 := big.NewInt(1)
	f2 := big.NewInt(1)

	fib <- f1; fib <- f2

	tmp := big.NewInt(0)
	for {
		tmp = tmp.Set(f1)
		f1 = f1.Add(f1, f2)
		f2 = f2.Set(tmp)

		fib <- big.NewInt(0).Set(f1)
	}
}
