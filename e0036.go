package main

func main() {

	sum := 0
	for n := 1; n < 1e6; n += 2 {
		if pal(digits(n, 2)) &&
		   pal(digits(n, 10)) {
			sum += n
		}
	}

	println(sum)
}

func pal(digits []int) bool {

	dlen := len(digits)
	for i := 0; i < dlen / 2; i++ {
		if digits[i] != digits[dlen - i - 1] {
			return false
		}
	}

	return true
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
