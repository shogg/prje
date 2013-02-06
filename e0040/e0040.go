package e0040

func E0040() int64 {
	result := 1
	for i := 1; i <= 1e6; i *= 10 {
		result *= d(i)
	}

	return int64(result)
}

func d(i int) int {

	b := 1; m, n := 1, 9
	for i > m*n {
		i -= m*n
		b *= 10; m += 1; n *= 10
	}

	number := b - 1 + (i + m - 1) / m
	digit := m - 1 - (i - 1) % m

	return digits(number, 10)[digit]
}

func digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n % base)
		n /= base
	}

	d = append(d, n)

	return d
}
