package e0054

import (
	"strings"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"github.com/shogg/prje"
)

var (
	highcard = []int { 0, 0, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192 }
)

const (
	A = 14
)

const (
	HIGHCARD = iota * 1e5
	PAIR
	TWO_PAIRS
	THREE
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR
	STRAIGHT_FLUSH
)

type card struct {
	val, color int
}

type hand []card

func E0054() int64 {

	hands := readHands("http://projecteuler.net/project/poker.txt")

	p1 := 0
	for _, h := range hands {
		s0 := score(h[0])
		s1 := score(h[1])
		if s0 > s1 {
			p1++
			fmt.Println(h[0], ">", h[1], "\t", s0, s1)
		} else {
			fmt.Println(h[0], "<", h[1], "\t", s0, s1)
		}
	}

	return int64(p1)
}

func score(h hand) int {

	str := straight(h)
	flu := flush(h)
	if str * flu > 0 {
		return STRAIGHT_FLUSH + str * flu
	}

	m := multi(h, 0, 1, 1)
	if m > 0 { return FULL_HOUSE + m }

	if flu > 0 { return FLUSH + flu }
	if str > 0 { return STRAIGHT + str }

	m = multi(h, 2, 0, 1)
	if m > 0 { return THREE + m }
	m = multi(h, 1, 2)
	if m > 0 { return TWO_PAIRS + m }
	m = multi(h, 3, 1)
	if m > 0 { return PAIR + m }

	s := 0
	for _, card := range h {
		s += highcard[card.val]
	}
	return s
}

func multi(h hand, counts ...int) int {
	m := 0
	for i, c := range count(h) {
		if c == 0 { continue }
		if c > len(counts) { return 0 }
		counts[c - 1]--
		if counts[c - 1] < 0 { return 0 }

		m += int(prje.Pow(16, c - 1)) * i
	}

	return m
}

func flush(h hand) int {
	last := h[0]
	for _, c := range h {
		if last.color != c.color { return 0 }
		last = c
	}

	max := 0
	for i, c := range count(h) {
		if c > 0 { max = i }
	}

	return max
}

func straight(h hand) int {
	counts := count(h)
	n := 0
	if counts[14] == 1 { n = 1 }
	for i, c := range counts {
		if i < 2 { continue }
		if c == 0 { n = 0; continue }
		n++
		if n == 5 {
			return i
		}
	}
	return 0
}

func count(h hand) []int {
	var counts = make([]int, 14 + 2)
	for _, c := range h {
		counts[c.val]++
	}
	return counts
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

func (h hand) String() string {
	if flush(h) > 0 {
		max := 0
		for i, c := range count(h) {
			if c > 0 { max = i }
		}
		return fmt.Sprint("flu_", card { max, 0 })
	}
	if straight(h) > 0 {
		max := 0
		for i, c := range count(h) {
			if c > 0 { max = i }
		}
		return fmt.Sprint("str_", card { max, 0 })
	}
	s := ""
	for i, c := range count(h) {
		for n := 0; n < c; n++ {
			s += fmt.Sprint(card { i, 0 })
		}
	}
	return s
}

func (c card) String() string {
	switch v := c.val; {
	case v == 10: return "T"
	case v == 11: return "J"
	case v == 12: return "Q"
	case v == 13: return "K"
	case v == 14: return "A"
	}

	return fmt.Sprint(c.val)
}
