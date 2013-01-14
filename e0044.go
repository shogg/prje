package main

import (
	"math"
)

func main() {

	d := int64(0)
	for i := int64(2); d == 0 ; i++ {
		for j := int64(i - 1); j > 0; j-- {
			p1 := pentagon(i) + pentagon(j)
			p2 := pentagon(i) - pentagon(j)
			if pentagonal(p1) && pentagonal(p2) {
				d = p2; break
			}
		}
	}

	println(d)
}

func pentagon(n int64) int64 {
	return n*(3*n - 1) / 2
}

func pentagonal(p int64) bool {
	n := 1.0/6.0 + math.Sqrt(1.0/36.0 + 2.0*float64(p)/3.0)
	return pentagon(int64(n)) == p
}
