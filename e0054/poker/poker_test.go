package poker

import "testing"

func TestConf(t *testing.T) {

	tests := []struct {
		hand         string
		handExpected string
		config       Conf
	}{
		{"1H 2C 3H 4D 5D", "5D 4D 3H 2C 1H", Conf{1, 1, 1, 1, 1}},
		{"4H 2C 3H 4D 5D", "4H 4D 5D 3H 2C", Conf{2, 2, 1, 1, 1}},
		{"4H 2C 5H 4D 5D", "5H 5D 4H 4D 2C", Conf{2, 2, 2, 2, 1}},
		{"4H 2C 3H 4D 4C", "4H 4D 4C 3H 2C", Conf{3, 3, 3, 1, 1}},
		{"4H 2C 2H 4D 4C", "4H 4D 4C 2C 2H", Conf{3, 3, 3, 2, 2}},
		{"4H 4S 2H 4D 4C", "4H 4S 4D 4C 2H", Conf{4, 4, 4, 4, 1}},
	}

	for _, tt := range tests {
		hand := ParseHand(tt.hand)
		handExpected := ParseHand(tt.handExpected)
		config := conf(hand)
		if config != tt.config {
			t.Errorf("%v erwartet, war %v", tt.config, config)
		}

		if *handExpected != *hand {
			t.Errorf("Hand '%s' erwartet, war '%s'", handExpected, hand)
		}
	}
}

func TestLess(t *testing.T) {

	tests := []struct {
		h1, h2 string
		win    bool
	}{
		{"5H 5C 6S 7S KD", "2C 3S 8S 8D TD", false},
		{"5D 8C 9S JS AC", "2C 5C 7D 8S QH", true},
		{"2D 9C AS AH AC", "3D 6D 7D TD QD", false},
		{"4D 6S 9H QH QC", "3D 6D 7H QD QS", true},
		{"2H 2D 4C 4D 4S", "3C 3D 3S 9S 9D", true},

		{"8C TS KC 9H 4S", "7D 2S 5D 3S AC", false},
		{"5C AD 5D AC 9C", "7C 5H 8D TD KS", true},
		{"3H 7H 6S KC JS", "QH TD JC 2D 8S", true},
		{"TH 8H 5C QS TC", "9H 4D JC KS JS", false},
		{"7C 5H KC QH JD", "AS KH 4C AD 4S", false},
	}

	for i, tt := range tests {
		h1 := ParseHand(tt.h1)
		h2 := ParseHand(tt.h2)
		if gt(h1, h2) != tt.win {
			t.Errorf("%d: '%s' > '%s' == %v erwartet", i, h1, h2, tt.win)
		}
	}
}

func gt(h1, h2 *Hand) bool {
	v1 := score(h1)
	v2 := score(h2)
	return v1 > v2
}
