package main

func main() {

	sum := 0
	for n := 3; n < 8 * factorial(9); n++ {

		s := 0
		for _, d := range digits(n) {
			s += factorial(d)
		}

		if s == n {
			sum += s
		}

	}

	println(sum)
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

func factorial(n int) int {

	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}

	return f
}
