package main

func main() {

	max := 0
	result := 0
	for p := 3; p <= 1000; p++ {
		count := 0
		for a := 1; a < p / 2; a++ {
			for b := a + 1; b < p - a; b++ {
				c := p - a - b
				if c*c == a*a + b*b {
					count++
				}
			}
		}

		if count > max { result = p; max = count }
	}

	println(result)
}
