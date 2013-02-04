package main

func main() {

	result := 0
	primes, sieve := psieve(999999)

out:
	for _, p := range primes {
		d := digits(p, 10)
		for min := 0; min <= 2; min++ {

			idx := indexes(d, min)
			for m := 1; log2(m) + 1 <= len(idx); m++ {

				count := 1
				d1 := make([]int, len(d)); copy(d1, d)
				for i := min; i < 9; i++ {
					n := number(inc(d1, mask(idx, m)))
					if sieve[n] == false {
						count++
					}
				}

				if count >= 8 {
					result = p; break out
				}
			}
		}
	}

	println(result)
}

func inc(digits, idx []int) []int {
	for _, i := range idx {
		digits[i]++
	}
	return digits
}

func mask(idx []int, m int) []int {

	var masked []int
	for _, i := range idx {
		if m & 0x1 == 1 {
			masked = append(masked, i)
		}
		m >>= 1
	}

	return masked
}

func indexes(digits []int, d int) []int {

	var idx []int
	for i, n := range digits {
		if n == d {
			idx = append(idx, i)
		}
	}

	return idx
}

func log2(n int) int {
	log := 0
	for ; n >= 2; n >>= 1 {
		log++
	}
	return log
}

func psieve(n int) ([]int, []bool) {

	var primes []int
	var sieve = make([]bool, n + 1)

	for p := 2; p <= n; p++ {
		if sieve[p] { continue }

		for i := 2*p; i <= n; i += p {
			sieve[i] = true
		}

		primes = append(primes, p)
	}

	return primes, sieve
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

func number(digits []int) int {

	n := 0
	for i, d := range digits {
		n += pow(10, i) * d
	}

	return n
}

func pow(n, e int) int {

	p := 1
	for i := 1; i <= e; i++ {
		p *= n
	}

	return p
}
