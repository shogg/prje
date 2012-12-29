package main

func main() {

	sum := 1
	n := 1
	for m := 2; m < 1001; m += 2 {
		n += m; sum += n
		n += m; sum += n
		n += m; sum += n
		n += m; sum += n
	}

	println(sum)
}
