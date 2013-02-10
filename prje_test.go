package prje_test

import (
	"testing"
	"github.com/shogg/prje"
)

var results = [1000]int64 {}

func init() {
		results[8] = int64(40824)
		results[9] = int64(31875000)
		results[10] = int64(142913828922)
		results[11] = int64(70600674)
		results[12] = int64(76576500)
		results[13] = int64(5537376230)
//		results[14] = int64()
//		results[15] = int64()
		results[16] = int64(1366)
		results[17] = int64(21124)
		results[18] = int64(1074)
		results[19] = int64(171)
		results[20] = int64(648)
		results[21] = int64(31626)
		results[22] = int64(871198282)
		results[23] = int64(4179871)
		results[24] = int64(2783915460)
		results[25] = int64(4782)
		results[26] = int64(983)
		results[27] = int64(-59231)
		results[28] = int64(669171001)
		results[29] = int64(9183)
		results[30] = int64(443839)
		results[31] = int64(73682)
		results[32] = int64(45228)
		results[33] = int64(100)
		results[34] = int64(40730)
		results[35] = int64(55)
		results[36] = int64(872187)
		results[37] = int64(748317)
		results[38] = int64(932718654)
		results[39] = int64(840)
		results[40] = int64(210)
		results[41] = int64(7652413)
		results[42] = int64(162)
		results[43] = int64(16695334890)
		results[44] = int64(5482660)
		results[45] = int64(1533776805)
		results[46] = int64(5777)
		results[47] = int64(134043)
		results[48] = int64(9110846700)
		results[49] = int64(296962999629)
		results[50] = int64(997651)
		results[51] = int64(121313)
		results[52] = int64(142857)
		results[53] = int64(4075)
		results[54] = int64(376)
}

func TestPrje(t *testing.T) {
	for i, p := range prje.Problems {
		if p == nil { continue }
		r := results[i]

		if r == 0 {
			t.Error("solution for problem", i, "is missing")
			continue
		}

		if n := p(); n != r {
			t.Error("problem", i, "expected", r, "was", n)
		}
	}
}
