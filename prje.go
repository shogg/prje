package prje

import (
	"github.com/shogg/prje/e0008"
	"github.com/shogg/prje/e0009"
	"github.com/shogg/prje/e0010"
	"github.com/shogg/prje/e0011"
	"github.com/shogg/prje/e0012"
	"github.com/shogg/prje/e0013"
	"github.com/shogg/prje/e0036"
)

var (
	Problems = [1000]func() int64 {}
)

func init() {
	Problems[8] = e0008.E0008
	Problems[9] = e0009.E0009
	Problems[10] = e0010.E0010
	Problems[11] = e0011.E0011
	Problems[12] = e0012.E0012
	Problems[13] = e0013.E0013
	Problems[36] = e0036.E0036
}
