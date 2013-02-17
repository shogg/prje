package e0024

import (
	"container/list"
	"github.com/shogg/prje"
)

func E0024() int64 {

	// faculty numbers descending
	f := make([]int, 9)
	f[8] = 1
	for n := 2; n <= 9; n++ {
		f[9-n] = n * f[9-n+1]
	}

	// digits 0..9
	d := list.New()
	for i := 0; i <= 9; i++ {
		d.PushBack(i)
	}

	p := make([]int, 0, 10)
	fsum := 0
	for i := 0; i < len(f); i++ {
		e := d.Front()
		for ; fsum + f[i] < 1e6; {
			fsum += f[i]
			e = e.Next()
		}

		d.Remove(e)
		p = append(p, e.Value.(int))
	}

	p = append(p, d.Front().Value.(int))
	p = prje.Reverse(p)
	return prje.Number(p, 10)
}
