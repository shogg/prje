package e0050

const N = 1e6

func E0050() int64 {

	primes, sieve := psieve(N)

	max := 0
	prime := 0

	for i, _ := range primes {
		sum := 0
		count := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			count++

			if sum > N { break }

			if !sieve[sum] && count > max {
				prime = sum
				max = count
			}
		}
	}

	return int64(prime)
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
