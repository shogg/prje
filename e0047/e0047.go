package e0047

import (
	"github.com/shogg/prje"
)

func E0047() int64 {

	result := 0
	for n := 2; ; n++ {

		ok := true
		for m := n; m < n + 4; m++ {
			if len(prje.Uniq(prje.Primefactors(m))) < 4 {
				ok = false; n = m; break
			}
		}

		if ok {
			result = n; break
		}
	}

	return int64(result)
}
