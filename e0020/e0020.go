package e0020

func E0020() int64 {

	e6 := factorial(100)

	D := make(chan int)
	go digits(e6, D)

	sum := 0
	for d := range D {
		sum += d
	}

	return int64(sum)
}

func factorial(n int) []int {

	e6 := []int { 1 }
	for i := 1; i <=n; i++ {
		u := 0
		for j := 0; j < len(e6); j++ {
			e := e6[j] * i + u
			u = 0
			if e >= 1e6 {
				u = e / 1e6; e %= 1e6
			}
			e6[j] = e
		}

		if u > 0 {
			e6 = append(e6, u)
		}
	}
	return e6
}

func digits(e6 []int, digits chan int) {

	for j := 0; j < len(e6); j++ {
		for e := e6[j]; e > 0; e /= 10 {
			d := e % 10
			digits <- d
		}
	}
	close(digits)
}
