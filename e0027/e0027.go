package e0027

func E0027() int64 {
	max, prod := 0, 0
	for a := -999; a <= 999; a++ {
		for b := -999; b <= 999; b++ {
			f, n := fn(a, b), 0
			for prime(f(n)) { n++ }
			if n > max { prod = a*b; max = n }
		}
	}

	return int64(prod)
}

func fn(a, b int) func(int) int {
	return func(n int) int { return n*n + a*n + b }
}

func prime(p int) bool {
	if p < 2 { return false }
	for i := 2; i < p; i++ {
		if p % i == 0 { return false }
	}

	return true
}
