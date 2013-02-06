package e0031

func E0031() int64 {
	pence := []int { 200, 100, 50, 20, 10, 5, 2, 1 }
	coins := []int { 1, 0, 0, 0, 0, 0, 0, 0 }

	sum := count(coins, pence) + 1
	return int64(sum)
}

func count(coins, pence []int) int {
	if len(coins) == 1 { return 0 }

	sum := 0
	for i := 1; i < len(coins) - 1 ; i++ {
		if coins[i] != 0 {
			sum += count(coins[i:], pence[i:])
			break
		}
	}

	for c := 1; c <= coins[0]; c++ {
		s := split(c * pence[0], pence[1:])
		s = plus(s, coins[1:])
		sum += count(s, pence[1:]) + 1
	}

	return sum
}

func plus(a, b []int) []int {
	for i := 0; i < len(a); i++ {
		a[i] += b[i]
	}
	return a
}

func split(value int, pence []int) []int {

	s := make([]int, len(pence))
	for i := 0; i < len(pence); i++ {
		s[i], value = value / pence[i], value % pence[i]
		if value == 0 { break }
	}

	return s
}
