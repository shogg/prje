package prje

func Digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n % base)
		n /= base
	}

	d = append(d, n)

	return d
}

func Number(digits []int, base int) int64 {

	n := int64(0)
	for i, d := range digits {
		n += Pow(base, i) * int64(d)
	}

	return n
}

func Pow(n, e int) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
	}

	return p
}

func Reverse(digits []int) []int {
	r := make([]int, len(digits))
	copy(r, digits)

	size := len(r)
	for i := range r {
		r[i], r[size - i - 1] = r[size - i - 1], r[i]
		if i >= size - i - 2 { break }
	}

	return r
}

func Pal(digits []int) bool {

	dlen := len(digits)
	for i := 0; i < dlen / 2; i++ {
		if digits[i] != digits[dlen - i - 1] {
			return false
		}
	}

	return true
}
