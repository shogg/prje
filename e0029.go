package main

import (
	"fmt"
	"sort"
)

const (
	N = 10
)

var (
	matrix [N+1][N+1]bool
)

type primefactor struct {
	prime, exp int
}

func main() {

	// 2^4 == 4^2 (2^(2*2))
	// 2^6 (2^(2*3)) == 8^2 (2^3^2) == 4^3 (2^2^3)
	// 2^8 == 4^4 (2^(2*4)) == 16^2 (2^(4*2))
	// 2^24 == (2^(3*2^3)) (2^(3*8)) (2^(4*6)) (2^(2*12))
	// 3^4 (3^(2*2)) == 9^2 (3^2^2)
	// 3^6 (3^(2*3)) == 27^2 (3^3^2) == 9^3 (3^2^3)
	// 4^6 (2^2^(2*3)) == 8^4 (2^3^(2*2)) (2^(2*2*3)) 
	// 5^4 == 25^2 (5^(2*2))
	// 5^6 == 25^4 (5^(2*3)) == 125^2 (5^(3*2))

	for a := 2; a <= N; a++ {
		factorsA := primefactors(a)
		for b := 2; b <= N; b++ {
			for _, fb := range factors(b) {
				a1 := pow(a, fb)
				b1 := b/fb
				if a1 <= N {
					matrix[a1][b1] = true
				}
				if len(factorsA) == 1 && factorsA[0].exp > 1 && factorsA[0].exp < fb {
					a2 := pow(factorsA[0].prime, fb)
					b2 := b/fb*factorsA[0].exp
					if a2 <= N && b2 <= N {
						matrix[a2][b2] = true
					}
				}
			}
		}
	}

	count := 0
	for a := 2; a <= N; a++ {
		for b := 2; b <= N; b++ {
			if matrix[a][b] == false {
				count++
			} else {
				fmt.Printf("%d^%d ", a, b)
			}
		}
	}
	fmt.Println()

	for a := 2; a <= N; a++ {
		fmt.Printf("%3d ", a)
		for b := 2; b <= N; b++ {
			if matrix[a][b] == true {
				fmt.Print("T "); continue
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}

	println(count)
	println(len(distinctPowers(N)))
}

func factors(n int) []int {

	var result []int
	for f := 2; f <= n; f++ {
		for n % f == 0 {
			result = append(result, f)
			n /= f
		}
	}

	return result
}

func primefactors(n int) []primefactor {

	var result []primefactor
	for f := 2; f <= n; f++ {
		pf := primefactor { f, 0 }
		for n % f == 0 {
			pf.exp++
			n /= f
		}

		if pf.exp > 0 {
			result = append(result, pf)
		}
	}

	return result
}

func pow(n, e int) int {

	p := 1
	for i := 0; i < e; i++ {
		p *= n
		if p > N { break }
	}

	return p
}

func powExact(n, e int) int {
	p := 1
	for i := 0; i < e; i++ {
		p *= n
	}

	return p
}

func distinctPowers(n int) []int {

	var powers []int
	for a := 2; a <= n; a++ {
		for b := 2; b <= n; b++ {
			powers = append(powers, powExact(a, b))
		}
	}

	sort.Ints(powers)

	var distinct []int
	last := 0
	for _, p := range powers {
		if p != last {
			distinct = append(distinct, p)
			last = p
		} else {
			println(p)
		}
	}

	return distinct
}
