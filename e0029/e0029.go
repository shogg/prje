package e0029

const N = 100

func E0029() int64 {

	var matrix [N+1][N+1]bool

	for a := 2; a <= N; a++ {
		prime, exp := same(primefactors(a))
		for b := 2; b <= N; b++ {
			a0, b0 := a, b
			if exp >= 2 {
				a0, b0 = prime, b * exp
			}
			for _, f := range factors(b0) {
				a1 := pow(a0, f)
				b1 := b0 / f
				if a1 > a && a1 <= N && b1 <= N {
					matrix[a1][b1] = true
				}
			}
		}
	}

	count := 0
	for a := 2; a <= N; a++ {
		for b := 2; b <= N; b++ {
			if ! matrix[a][b] {
				count++
			}
		}
	}

	return int64(count)
}

func factors(n int) []int {

	var result []int
	for f := 2; f < n; f++ {
		if n % f == 0 {
			result = append(result, f)
		}
	}

	return result
}

func primefactors(n int) []int {

	var result []int
	for f := 2; f <= n; f++ {
		for n % f == 0 {
			result = append(result, f)
			n /= f
		}
	}

	return result
}

func same(f []int) (prime, exp int) {

	cmp := f[0]
	for _, p := range f {
		if p != cmp { return 0, 0 }
	}

	return cmp, len(f)
}

func pow(n, e int) int {

	p := 1
	for i := 0; i < e; i++ {
		p *= n
		if p > N { break }
	}

	return p
}
