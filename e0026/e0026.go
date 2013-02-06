package e0026

func E0026() int64 {

	cmax := 0
	dmax := 0
	for d := 1; d <= 1000; d++ {
		c := cycle(1, d)
		if c > cmax { cmax = c; dmax = d }
	}

	return int64(dmax)
}

func cycle(n, d int) int {

	for n < d { n *= 10 }
	n %= d

	M := []int { n }

	c := 1
	for {
		n = n * 10 % d
		i := index(M, n)
		if i >= 0 { c -= i; break }

		M = append(M, n)
		c++
	}

	return c
}

func index(A []int, e int) int {
	for i, a := range A {
		if a == e { return i }
	}

	return -1
}
