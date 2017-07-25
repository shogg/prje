package poker

import "sort"
import "strings"
import "strconv"
import "bytes"

type Color byte

const (
	Heart   Color = 'H'
	Diamond Color = 'D'
	Spade   Color = 'S'
	Club    Color = 'C'
)

type Card struct {
	Rank uint
	Color
}

type Hand [5]Card

const (
	HighCard uint = iota << 28
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type Conf [5]uint8

var (
	confFourOfAKind  = Conf{4, 4, 4, 4, 1}
	confFullHouse    = Conf{3, 3, 3, 2, 2}
	confThreeOfAKind = Conf{3, 3, 3, 1, 1}
	confTwoPairs     = Conf{2, 2, 2, 2, 1}
	confOnePair      = Conf{2, 2, 1, 1, 1}
	confHighCard     = Conf{1, 1, 1, 1, 1}
)

// Magnitudes
const (
	m1 = 14
	m2 = m1 * 14
	m3 = m2 * 14
	m4 = m3 * 14
)

func score(h *Hand) uint {

	var s uint
	conf := conf(h)

	// Copy to allow different orderings.
	hcopy := *h

	switch {
	case straightFlush(&hcopy):
		s = StraightFlush + hcopy[0].Rank
	case straight(&hcopy):
		s = Straight + hcopy[0].Rank
	case flush(&hcopy):
		h4 := hcopy[0].Rank
		h3 := hcopy[1].Rank
		h2 := hcopy[2].Rank
		h1 := hcopy[3].Rank
		h0 := hcopy[4].Rank
		s = Flush + h4*m4 + h3*m3 + h2*m2 + h1*m1 + h0
	case confFourOfAKind == conf:
		s = FourOfAKind + h[0].Rank
	case confFullHouse == conf:
		s = FullHouse + h[0].Rank
	case confThreeOfAKind == conf:
		s = ThreeOfAKind + h[0].Rank
	case confTwoPairs == conf:
		h2 := h[0].Rank
		h1 := h[2].Rank
		h0 := h[4].Rank
		s = TwoPairs + h2*m2 + h1*m1 + h0
	case confOnePair == conf:
		h3 := h[0].Rank
		h2 := h[2].Rank
		h1 := h[3].Rank
		h0 := h[4].Rank
		s = OnePair + h3*m3 + h2*m2 + h1*m1 + h0
	case confHighCard == conf:
		h4 := h[0].Rank
		h3 := h[1].Rank
		h2 := h[2].Rank
		h1 := h[3].Rank
		h0 := h[4].Rank
		s = HighCard + h4*m4 + h3*m3 + h2*m2 + h1*m1 + h0
	}

	return s
}

func straightFlush(h *Hand) bool {
	return straight(h) && flush(h)
}

func straight(h *Hand) bool {
	sort.Stable(byRank{h})

	for i := range h {
		if i == 0 {
			continue
		}
		if h[i-1].Rank-1 != h[i].Rank {
			return false
		}
	}

	return true
}

func flush(h *Hand) bool {
	sort.Stable(byRank{h})

	for i := range h {
		if h[0].Color != h[i].Color {
			return false
		}
	}

	return true
}

func conf(h *Hand) Conf {
	sort.Stable(byRank{h})

	m := make(map[int][]Card)

	var last Card
	count := 1
	for i := range h {

		if i == 0 {
			last = h[i]
			continue
		}

		if h[i].Rank == last.Rank {
			count++
			continue
		}

		for j := count; j >= 1; j-- {
			m[count] = append(m[count], h[i-j])
		}

		last = h[i]
		count = 1
	}

	for j := count; j >= 1; j-- {
		m[count] = append(m[count], h[5-j])
	}

	var c Conf
	index := 0
	for i := 4; i >= 1; i-- {
		for j := 0; j < len(m[i]); j++ {
			h[index] = m[i][j]
			c[index] = uint8(i)
			index++
		}
	}

	return c
}

type byRank struct {
	*Hand
}

func (s byRank) Len() int {
	return len(s.Hand)
}

func (s byRank) Less(i, j int) bool {
	return -s.Hand[i].Rank < -s.Hand[j].Rank
}

func (s byRank) Swap(i, j int) {
	s.Hand[i], s.Hand[j] = s.Hand[j], s.Hand[i]
}

func (h1 *Hand) Less(h2 *Hand) bool {
	return score(h1) < score(h2)
}

func ParseHand(s string) *Hand {
	var h Hand
	comp := strings.Split(strings.TrimSpace(s), " ")
	for i := 0; i < 5; i++ {
		h[i] = ParseCard(comp[i])
	}

	return &h
}

func ParseCard(s string) Card {

	rank := 0
	switch s[0] {
	case 'A':
		rank = 14
	case 'K':
		rank = 13
	case 'Q':
		rank = 12
	case 'J':
		rank = 11
	case 'T':
		rank = 10
	default:
		rank, _ = strconv.Atoi(string(s[0]))
	}

	return Card{
		Rank:  uint(rank),
		Color: Color(s[1]),
	}
}

func (h *Hand) String() string {

	var buf bytes.Buffer
	for i := range h {
		if i > 0 {
			buf.WriteString(" ")
		}

		buf.WriteString(h[i].String())
	}

	return buf.String()
}

var above9 = [5]string{"T", "J", "Q", "K", "A"}

func (c Card) String() string {
	r := strconv.Itoa(int(c.Rank))
	if c.Rank > 9 {
		r = above9[c.Rank-10]
	}
	return r + string(c.Color)
}
