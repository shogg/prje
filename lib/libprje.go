package libprje

func Digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n % base)
		n /= base
	}

	d = append(d, n)

	return d
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
