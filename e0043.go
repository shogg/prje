package main

const (
	N = 10
)

func main() {

	digits := [...]int { 9, 8, 7, 6, 5, 4, 3, 2, 1, 0 }
	primes := []int { 2, 3, 5, 7, 11, 13, 17 }

	sum := int64(0)
	for i := 0; i < factorial(N); i++ {
		p := permutation(i, digits)

		ok := true
		for j := 0; j < len(primes); j++ {
			n3 := number(p[N-j-4:N-j-1])
			prime := int64(primes[j])
			ok = n3 % prime == 0

			if !ok { break }
		}

		if ok {
			sum += number(p)
		}
	}

	println(sum)
}

func permutation(i int, digits [N]int) []int {

	p := make([]int, len(digits))
	f := factorialDigits(i)
	for i := 0; i < len(f); i++ {
		fi := f[N-i-1]
		p[N-i-1] = digits[fi]
		if fi < len(digits) - i - 1 {
			copy(digits[fi:], digits[fi + 1:])
		}
	}

	return p
}

func factorialDigits(n int) []int {

	f := make([]int, N)
	for i := 1; n != 0; i++ {
		f[i - 1] = n % i
		n /= i
	}

	return f
}

func factorial(n int) int {

	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}

	return f
}

func number(digits []int) int64 {

	n := int64(0)
	for i, d := range digits {
		n += pow(10, i) * int64(d)
	}

	return n
}

func pow(n, e int) int64 {

	p := int64(1)
	for i := 1; i <= e; i++ {
		p *= int64(n)
	}

	return p
}
