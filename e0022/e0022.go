package e0022

import (
	"net/http"
	"io/ioutil"
	"strings"
	"sort"
)

func E0022() int64 {

	resp, err := http.Get("http://projecteuler.net/project/names.txt")
	if err != nil { panic(err) }

	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	N := names(string(text))
	sort.Strings(N)

	sum := 0
	for i, n := range N {
		sum += (i+1) * score(n)
	}

	return int64(sum)
}

func names(s string) []string {

	n := strings.Split(s, ",")
	for i := 0; i < len(n); i++ {
		n[i] = strings.Trim(n[i], "\"\t ")
	}

	return n
}

func score(name string) int {

	s := 0
	for _, n := range name {
		s += int(n) - int('A') + 1
	}

	return s
}
