package e0048

func E0048() int64 {

	sum := int64(0)
	for a := 1; a <= 1e3; a++ {
		sum += powMod(a, a, 1e10)
	}

	sum %= 1e10
	return sum
}

func powMod(n, e int, mod int64) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
		p %= mod
	}

	return p
}
