package main

import (
	"sort"
)

func main() {

	n := 1
	for ; ; n++ {
		d1 := digits(n, 10)
		sort.Ints(d1)
		found := true
		for x := 2; x <= 6; x++ {
			dx := digits(n*x, 10)
			sort.Ints(dx)
			if !sameset(d1, dx) {
				found = false; break
			}
		}
		if found { break }
	}

	println(n)
}

func sameset(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] { return false }
	}
	return true
}

func digits(n, base int) []int {

	var d []int
	for n >= base {
		d = append(d, n % base)
		n /= base
	}

	d = append(d, n)

	return d
}
