package e0030

func E0030() int64 {

	var numbers []int

	for n := 10; n <= 5 * pow(9, 5); n++ {
		psum := 0
		for _, d := range digits(n) {
			psum += pow(d, 5)
		}

		if psum == n {
			numbers = append(numbers, n)
		}
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return int64(sum)
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

func pow(n, e int) int {

	p := n
	for i := 1; i < e; i++ {
		p *= n
	}

	return p
}
