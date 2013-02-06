package e0016

func E0016() int64 {

	digits := make([]int, 1, 302)
	digits[0] = 1
	for n := 0; n < 1000; n++ {
		u := 0
		for i, d := range digits {
			 d = d * 2 + u; u = 0
			 if d > 9 { u = d / 10; d %= 10 }
			 digits[i] = d
		}

		if u > 0 {
			m := len(digits)
			digits = digits[0:m+1]
			digits[m] = u
		}

	}

	sum := 0
	for _, d := range digits {
		sum += d
	}

	return int64(sum)
}
