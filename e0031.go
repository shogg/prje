package main

import (
	"fmt"
)

// 73652 - wrong

func main() {
	pence := []int { 200, 100, 50, 20, 10, 5, 2, 1 }
	coins := []int { 0, 0, 1, 0, 0, 0, 0, 0 }

	head(pence)
	debug(coins)
	sum := count(1, coins, pence)
	println(sum)
}

func count(i int, coins, pence []int) int {
	if i >= len(coins) - 1 { return 1 }

	sum := 1
	for c := 1; c <= coins[i]; c++ {
		for j := i + 1; j < len(coins) - 1; j++ {
			sum += count(j, coins, pence)
		}
		s := split(i+1, c * pence[i], pence)
		s = plus(s, coins)
		s[i] = coins[i] - c
		debug(s)
		sum += count(i+1, s, pence)
	}

	return sum
}

func plus(a, b []int) []int {
	for i := 0; i < len(a); i++ {
		a[i] += b[i]
	}
	return a
}

func split(i, value int, pence []int) []int {

	s := make([]int, len(pence))
	for ; i < len(pence); i++ {
		s[i], value = value / pence[i], value % pence[i]
		if value == 0 { break }
	}

	return s
}

func head(pence []int) {
	for i := 7; i >= 0; i-- {
		print(fmt.Sprintf("%3d ", pence[len(pence) - i - 1]))
	}
	println(fmt.Sprint("\n-------------------------------"))
}

func debug(coins []int) {
	for i := 7; i >= 0; i-- {
		if i > len(coins) - 1 || coins[len(coins) - i - 1] == 0 {
			fmt.Print("    ")
		} else {
			fmt.Printf("%3d ", coins[len(coins) - i - 1])
		}
	}
	fmt.Println()
}
