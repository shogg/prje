package e0017

import (
	"strconv"
	"strings"
)

var (
	oner = []string {
		"", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine" }
	teens = []string {
		"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen" }
	tys = []string {
		"", "", "twenty", "thirty", "forty",
		"fifty", "sixty", "seventy", "eighty", "ninety" }
	magnitude = []string { "", "", "hundred", "thousand" }
)

func E0017() int64 {

	chars := 0
	for n := 1; n <= 1000; n++ {
		d := split(n)
		w := written(d)
		chars += len(w)
	}

	return int64(chars)
}

func split(n int) []int {
	str := strings.Split(strconv.Itoa(n), "")
	d := make([]int, len(str))
	for i, s := range str {
		d[len(d)-i-1], _ = strconv.Atoi(s)
	}

	return d
}

func written(d []int) string {

	s := ""
	for i := 0; i < len(d); i++ {
		if d[i] == 0 { continue }
		if len(d) >= 2 && i == 0 && d[1] == 1 { continue }
		if i == 1 && d[1] == 1 { s += teens[d[0]]; continue }
		if i == 1 { s += tys[d[i]]; continue }
		s += oner[d[i]] + magnitude[i]
	}

	if len(d) > 2 && (d[0] != 0 || d[1] != 0) {
		s += "and"
	}

	return s
}
