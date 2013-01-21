package main

import (
	"strings"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

var (
	highcard = []int { 0, 0, 2, 3, 6, 12, 24, 46, 89, 172, 332, 640, 1234, 2379, 4586 }
)

type card struct {
	val, color int
}

type hand []card

func main() {

	hands := readHands("http://projecteuler.net/project/poker.txt")

	p1 := 0
	for _, h := range hands {
		s0 := score(h[0]); print("\t")
		s1 := score(h[1])
		if s0 > s1 {
			p1++
		}
		println()
	}

	println(p1)
}

func score(h hand) int64 {
	str := straight(h); fl := flush(h)
	s := int64(multi(h) + str + fl)
	s += int64(str) * int64(fl)
	print(s)
	return s
}

func flush(h hand) int {
	last := h[0]
	for _, c := range h {
		if last.color != c.color { return 0 }
		last = c
	}

	max := 0
	for i, c := range count(h) {
		if c != 0 { max = i }
	}

	print("F ")
	return 3500 + max
}

func straight(h hand) int {
	counts := count(h)
	n := 0
	for i, c := range counts {
		if c == 0 { n = 0; continue }
		n++
		if n == 5 {
			print("S ")
			return 3400 + i
		}
		if n == 4 && counts[2] == 1 && counts[14] == 1 {
			print("S ")
			return 3400 + i
		}
	}
	return 0
}

func multi(h hand) int {
	val := 0
	for i, c := range count(h) {
		if c == 0 { continue }
		val += pow(15, c) * i
		switch c {
		case 1: print("   ")
		case 2: print("P2 ")
		case 3: print("K3 ")
		case 4: print("K4 ")
		}
	}

	return val
}

func count(h hand) []int {
	var counts = make([]int, 14 + 2)
	for _, c := range h {
		counts[c.val]++
	}
	return counts
}

func pow(n, e int) int {

	p := 1
	for i := 1; i <= e; i++ {
		p *= n
	}

	return p
}

func readHands(url string) [][2]hand {

	t0 := time.Now()
	poker_txt := read(url)
	fmt.Println("GET", time.Now().Sub(t0))

	var hands [][2]hand

	rows := strings.Split(poker_txt, "\n")
	for _, r := range rows {
		vals := strings.Split(r, " ")
		if len(vals) != 10 { continue }

		var h2 [2]hand
		for i := 0; i <= 1; i++ {
			h2[i] = parseHand(vals[i*5:(i*5)+5])
		}

		hands = append(hands, h2)
	}

	return hands
}

func read(url string) string {

	resp, err := http.Get(url)
	if err != nil { panic(err) }

	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	return string(text)
}

func parseHand(s []string) hand {
	h := make([]card, 5)
	for i := 0; i < 5; i++ {
		h[i] = parseCard(s[i])
	}
	return h
}

func parseCard(s string) card {

	val := 0
	switch s[0] {
	case 'T': val = 10
	case 'J': val = 11
	case 'Q': val = 12
	case 'K': val = 13
	case 'A': val = 14
	default: val = int(s[0]) - int('0')
	}
	c := card { val, int(s[1]) }
	return c
}
