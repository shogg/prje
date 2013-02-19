package e0042

import (
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/shogg/prje"
)

func E0042() int64 {

	words := read("http://projecteuler.net/project/words.txt")

	count := int64(0)
	for _, w := range words {
		if prje.Triangular(number(w)) {
			count++
		}
	}

	return count
}

func number(word string) int {
	n := 0
	for _, w := range word {
		n += int(w) - int('A') + 1
	}
	return n
}

func read(url string) []string {

	resp, err := http.Get(url)
	if err != nil { panic(err) }

	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	n := strings.Split(string(text), ",")
	for i := 0; i < len(n); i++ {
		n[i] = strings.Trim(n[i], "\"\t ")
	}

	return n
}
