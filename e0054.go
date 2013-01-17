package main

import (
	"strings"
//	"fmt"
	"net/http"
	"io/ioutil"
)

type card struct {
	val, color int
}

type hand []card

func main() {

	hands := readHands("")

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
	s := int64(multiscore(h) + straight(h) + flush(h))
	s += int64(straight(h)) * int64(flush(h))
	print(s)
	return s
}

func flush(h hand) int {
	last := h[0]
	for _, c := range h {
		if last.color != c.color { return 0 }
		last = c
	}
	print("f")
	return 6e4 + h[4].val
}

func straight(h hand) int {
	counts := count(h)
	n := 0
	for i, c := range counts {
		if c == 0 { n = 0; continue }
		n++
		if n == 5 {
			print("s")
			return 5e4 + i
		}
	}
	return 0
}

func multiscore(h hand) int {
	counts := count(h)
	val := 0
	for i, c := range counts {
		if c == 0 { continue }
		val += pow(15, c) + i
		switch c {
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
	poker_txt := read("http://projecteuler.net/project/poker.txt")

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

	println("HTTP ERFOLGREICH --------------------")
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
