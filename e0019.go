package main

var (
	months = []int { 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 }
)

func main() {

	sundays := 0

	y := 1; m := 0; d := -1
	for ; y < 101; {
		dm := months[m]
		if m == 1 && y % 4 == 0 { dm++ }
		if d > dm { d -= dm; m++ }
		if d % dm == 1 { sundays++ }
		if m > 11 { m = 0; y++ }

		d += 7
	}

	println(sundays)
}
