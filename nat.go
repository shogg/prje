package prje

import (
	"strconv"
)

type Nat struct {
	Digits []int
	Base int
}

func NewNat(digits []int, base int) Nat {
	return Nat { digits, base }
}

func (a Nat) Plus(b Nat) Nat {

	max := Max(len(a.Digits), len(b.Digits))
	sum := Nat { make([]int, max, max + 1), a.Base }

	carry := 0
	for i := 0; i < max; i++ {
		da := 0; if i < len(a.Digits) { da = a.Digits[i] }
		db := 0; if i < len(b.Digits) { db = b.Digits[i] }
		sum.Digits[i] = da + db + carry
		carry = sum.Digits[i] / a.Base
		sum.Digits[i] %= a.Base
	}

	if carry > 0 {
		sum.Digits = append(sum.Digits, carry)
	}

	return sum
}

func (nat Nat) String() string {
	chars := make([]uint8, len(nat.Digits))


	for i, n := range Reverse(nat.Digits) {
		chars[i] = strconv.Itoa(n)[0]
	}

	return string(chars)
}
