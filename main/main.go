package main

import (
	"flag"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/shogg/prje/problems"
)

func main() {

	cpuprof := flag.Bool("cpuprof", false, "aktiviert cpu-profiling")
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

	if *cpuprof {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			return
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			return
		}
		defer pprof.StopCPUProfile()
	}

	problem := problems.Problems[nr]
	if problem == nil {
		println("problem", nr, "not found")
		return
	}

	println(problem())
}
