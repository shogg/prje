package e0021

func E0021() int64 {

	sum := 0
	for a := 1; a < 10000; a++ {
		if b := d(a); d(b) == a && a != b {
			sum += a
		}
	}

	return int64(sum)
}

func d(n int) int {

	div := make(chan int)
	go divisors(n, div)

	sum := 0
	for i := range div {
		sum += i
	}

	return sum
}

func divisors(n int, div chan int) {

	for i := 1; i < n; i++ {
		if n % i != 0 { continue }
		div <- i
	}
	close(div)
}
