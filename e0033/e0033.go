package e0033

type rat struct {
	n, d int
}

func E0033() int64 {

	var list []rat
	for a := 11; a <= 98; a++ {
		digitsa := digits(a)
		for b := a + 1; b <= 99; b++ {
			if a == b { continue }
			r := rat { a, b }
			for _, d := range digitsa {
				if d == 0 { continue }
				a1 := cancel(digitsa, d)
				b1 := cancel(digits(b), d)
				if r.equals(rat { a1, b1 }) {
					list = append(list, r)
				}
			}
		}
	}

	result := rat { 1, 1 }
	for _, r := range list {
		result = result.times(r)
	}

	g := gcd(result.n, result.d)
	result = result.divide(rat { g, g })

	return int64(result.d)
}

func digits(n int) []int {

	var d []int
	for n >= 10 {
		d = append(d, n % 10)
		n /= 10
	}

	d = append(d, n)

	return d
}

func cancel(digits []int, d int) int {

	n := 0
	if digits[0] == d {
		n = digits[1]
	} else if digits[1] == d {
		n = digits[0]
	}

	return n
}

func (a rat) equals(b rat) bool {
	if a.d == 0 || b.d == 0 { return false }
	ra := float64(a.n) / float64(a.d)
	rb := float64(b.n) / float64(b.d)
	return ra == rb
}

func (a rat) times(b rat) rat {
	return rat { a.n * b.n, a.d * b.d }
}

func (a rat) divide(b rat) rat {
	return rat { a.n / b.n, a.d / b.d }
}

func gcd(a, b int) int {

	min := a
	if b < a { min = b }

	for n := min; n >= 1; n++ {
		if a % n == 0 && b % n == 0 {
			return n
		}
	}

	return 1
}
