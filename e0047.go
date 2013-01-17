package main

func main() {

	result := 0
	for n := 2; ; n++ {

		ok := true
		for m := n; m < n + 4; m++ {
			if len(primefactors(m)) < 4 {
				ok = false; n = m; break
			}
		}

		if ok {
			result = n; break
		}
	}

	println(result)
}

func primefactors(n int) []int {

	var result []int
	for f := 2; f <= n; f++ {
		if n % f == 0 {
			result = append(result, f)
		}
		for n % f == 0 {
			n /= f
		}
	}

	return result
}
