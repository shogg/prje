package e0028

func E0028() int64 {

	sum := 1
	n := 1
	for m := 2; m < 1001; m += 2 {
		n += m; sum += n
		n += m; sum += n
		n += m; sum += n
		n += m; sum += n
	}

	return int64(sum)
}
