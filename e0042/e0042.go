package e0042

import (
	"net/http"
	"io/ioutil"
	"strings"
	"math"
)

func E0042() int64 {

	text := read("http://projecteuler.net/project/words.txt")
	words := split(text)

	count := 0
	for _, w := range words {
		if triangle(number(w)) {
			count++
		}
	}

	return int64(count)
}

func triangle(t int) bool {
	n := -.5 + math.Sqrt(.25 + 2.0*float64(t))
	return n == math.Trunc(n)
}

func split(text string) []string {
	n := strings.Split(text, ",")
	for i := 0; i < len(n); i++ {
		n[i] = strings.Trim(n[i], "\"\t ")
	}

	return n
}

func number(word string) int {
	n := 0
	for _, w := range word {
		n += int(w) - int('A') + 1
	}
	return n
}

func read(url string) string {

	resp, err := http.Get(url)
	if err != nil { panic(err) }

	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	return string(text)
}
