package main

func main() {
	a, b, c := triplet()
	println(a * b * c)
}

func triplet() (a, b, c int) {
	for a := 1; a < 998; a++ {
		for b := a + 1; b < 999; b++ {
			for c := b + 1; c < 1000; c++ {
				if a*a + b*b == c*c && a + b + c == 1000 {
					return a, b, c
				}
			}
		}
	}
	return -1, -1, -1
}
