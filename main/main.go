package main

import (
	"flag"
	"strconv"
	"github.com/shogg/prje/problems"
)

func main() {

	flag.Parse()
	if flag.NArg() == 0 {
		println("usage: problem <nr>")
		return
	}

	arg0 := flag.Arg(0)
	nr, err := strconv.Atoi(arg0)
	if err != nil {
		println("usage: problem <nr>")
		return
	}

	problem := problems.Problems[nr]
	if problem == nil {
		println("problem", nr, "not found")
		return
	}

	println(problem())
}
