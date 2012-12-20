package main

func main() {

	primes := make(chan int)
	go psieve(2000000, primes)

	sum := int64(0)
	for p := range primes {
		sum += int64(p)
	}

	println(sum)
}

func psieve(n int, primes chan int) {

	sieve := make([]bool, n + 1)
	for p := 2; p <= n; p++ {
		if sieve[p] { continue }

		for i := 2*p; i <= n; i += p {
			sieve[i] = true
		}

		primes <- p
	}

	close(primes)
}
