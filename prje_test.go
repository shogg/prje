package prje_test

import (
	"testing"
	"github.com/shogg/prje"
)

var (
	results = [1000]int64 {}
)

func init() {
		results[36] = int64(872187)
}

func TestE0036(t *testing.T) {
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
